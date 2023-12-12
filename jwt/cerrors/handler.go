package cerrors

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Handler(ctx *fiber.Ctx, err error) error {
	fmt.Println("handler!")
	return nil
}
