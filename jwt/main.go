package main

import (
	"fmt"
	"goback/jwt/cerrors"
	"goback/jwt/dto"
	"goback/jwt/middleware"
	"goback/jwt/token"
	"goback/jwt/utils"

	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/golang-jwt/jwt/v5"
)

var users = make([]*dto.UserDto, 0)

func main() {

	app := fiber.New(fiber.Config{
		ErrorHandler: cerrors.Handler,
	})
	app.Use(logger.New())

	app.Post("/api/v1/signin", func(ctx *fiber.Ctx) error {
		userDto := new(dto.UserDto)
		if err := ctx.BodyParser(userDto); err != nil {
			return err
		}

		// ID, Password가 일치한지 확인
		isFind := false
		for _, user := range users {
			if *user == *userDto {
				isFind = true
				break
			}
		}
		if isFind == false {
			return utils.StatusBadRequest(ctx, &dto.GenericDto{
				Code:    fiber.StatusBadRequest,
				Message: "회원가입을 진행해주세요.",
			})
		}

		// Access-Token과 Refresh-Token 발급
		accessToken, err := token.Genreate(jwt.MapClaims{
			"iss": userDto.ID,
			"exp": time.Now().Add(time.Hour * 3).Unix(),
		})
		refreshToken, err := token.Genreate(jwt.MapClaims{
			"iss": userDto.ID,
			"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
		})
		if err != nil {
			return utils.StatusInternalServerError(ctx, &dto.GenericDto{
				Code:    fiber.StatusInternalServerError,
				Message: "토큰을 생성하지 못했습니다.",
			})
		}

		// 성공적으로 반환
		return ctx.Status(fiber.StatusCreated).JSON(&dto.GenericDto{
			Code:    fiber.StatusCreated,
			Message: "토큰을 정상적으로 생성했습니다.",
			Data: &dto.TokensDto{
				AccessToken:  accessToken,
				RefreshToken: refreshToken,
			},
		})
	})

	app.Post("/api/v1/signup", func(ctx *fiber.Ctx) error {
		userDto := new(dto.UserDto)
		if err := ctx.BodyParser(userDto); err != nil {
			return utils.StatusBadRequest(ctx, &dto.GenericDto{
				Code:    fiber.StatusBadRequest,
				Message: "body값의 요청 형식이 알맞지 않습니다.",
			})
		}

		// ID 중복 확인
		isFind := false
		for _, user := range users {
			if userDto.ID == user.ID {
				isFind = true
				break
			}
		}
		if isFind {
			return ctx.Status(fiber.StatusConflict).JSON(&dto.GenericDto{
				Code:    fiber.StatusConflict,
				Message: "이미 있는 ID입니다.",
			})
		}

		// 회원가입
		users = append(users, userDto)

		// 성공적으로 반환
		return utils.StatusBadRequest(ctx, &dto.GenericDto{
			Code:    fiber.StatusOK,
			Message: "로그인 성공",
		})
	})

	app.Post("/api/v1/jwt/refresh", func(ctx *fiber.Ctx) error {
		refreshTokenDto := new(dto.RefreshTokenDto)
		if err := ctx.BodyParser(refreshTokenDto); err != nil {
			return utils.StatusBadRequest(ctx, &dto.GenericDto{
				Code:    fiber.StatusBadRequest,
				Message: "body값의 요청 형식이 알맞지 않습니다.",
			})
		}

		mapClaims, err := token.Validate(refreshTokenDto.RefreshToken)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(&dto.GenericDto{
				Code:    fiber.StatusUnauthorized,
				Message: "RefreshToken이 유효하지 않습니다.",
			})
		}

		t, err := token.Genreate(jwt.MapClaims{
			"iss": mapClaims["iss"],
			"exp": time.Now().Add(time.Hour * 3).Unix(),
		})
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(&dto.GenericDto{
				Code:    fiber.StatusUnauthorized,
				Message: "토큰 생성시 오류가 발생했습니다.",
			})
		}

		return ctx.Status(fiber.StatusCreated).JSON(&dto.GenericDto{
			Code:    fiber.StatusCreated,
			Message: "Access token 재발급 완료",
			Data: &dto.RefreshTokenDto{
				RefreshToken: t,
			},
		})
	})

	app.Post("/api/v1/jwt/validation", middleware.Authenticate, func(ctx *fiber.Ctx) error {
		mapCliams, _ := ctx.Locals("validMapClaims").(jwt.MapClaims)
		issuer := mapCliams["iss"]

		return ctx.Status(fiber.StatusOK).JSON(&dto.GenericDto{
			Code:    fiber.StatusOK,
			Message: "Hi!" + fmt.Sprint(issuer),
		})
	})

	log.Fatal(app.Listen(":9190"))

}
