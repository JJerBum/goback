package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// !! 서버를 제작할 때는 이렇게 해서는 절대로 안됌!!
const SCERET_KEY = "1f8a0ca3b85e40f30fe32fe1a4fba64f881a8ab6d3df180159da865d9bf374e2973369326e2b2dc8476792e4cc986ecfbe32637710f50e0fc9aa754b505d04c6"

func Genreate(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(SCERET_KEY))
}

func Validate(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok == false {
			return nil, fmt.Errorf("전달된 토큰의 서명 알고리즘이 옳지 않습니다.: %v", token.Header["alg"])
		}

		return []byte(SCERET_KEY), nil
	})
	if err != nil {
		return nil, err
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	if token.Valid == false {
		return nil, fmt.Errorf("해당 토큰이 유효하지 않습니다.")
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		return nil, fmt.Errorf("토큰이 만료되었습니다.")
	}

	return claims, nil
}
