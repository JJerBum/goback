package utils

import "github.com/gofiber/fiber/v2"

func StatusBadRequest(ctx *fiber.Ctx, data any) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(data)
}

func StatusInternalServerError(ctx *fiber.Ctx, data any) error {
	return ctx.Status(fiber.StatusInternalServerError).JSON(data)
}
