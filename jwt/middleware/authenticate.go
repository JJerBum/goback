package middleware

import (
	"goback/jwt/dto"
	"goback/jwt/token"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Authenticate(ctx *fiber.Ctx) error {
	//토큰 추출
	tokenString := ctx.Get("Authorization")
	if tokenString == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&dto.GenericDto{
			Code:    fiber.StatusUnauthorized,
			Message: "토큰을 소지하지 않고 요청하였습니다.",
		})
	}
	tokenString = strings.Split(tokenString, " ")[1]

	mapClaims, err := token.Validate(tokenString)
	if err != nil {
		ctx.Status(fiber.StatusUnauthorized).JSON(&dto.GenericDto{
			Code:    fiber.StatusUnauthorized,
			Message: err.Error(),
		})
	}

	ctx.Locals("validMapClaims", mapClaims)

	return ctx.Next()
}
