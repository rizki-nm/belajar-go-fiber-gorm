package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizki-nm/belajar-go-fiber-gorm/model/web"
	"github.com/rizki-nm/belajar-go-fiber-gorm/utils"
)

func Auth(ctx *fiber.Ctx) error {
	token := ctx.Get("X-TOKEN")

	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(web.WebResponseFailed{
			Code:    fiber.StatusUnauthorized,
			Status:  "Unauthorized",
			Message: "User not authenticated",
		})
	}

	//_, err := utils.VerifyToken(token)

	claims, err := utils.DecodeToken(token)

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(web.WebResponseFailed{
			Code:    fiber.StatusUnauthorized,
			Status:  "Unauthorized",
			Message: "User not authenticated",
		})
	}

	role := claims["role"].(string)

	if role != "admin" {
		return ctx.Status(fiber.StatusForbidden).JSON(web.WebResponseFailed{
			Code:    fiber.StatusForbidden,
			Status:  "Forbidden",
			Message: "Forbidden access",
		})
	}

	// set
	//ctx.Locals("userInfo", claims)
	// access
	//ctx.Locals("userInfo")

	return ctx.Next()
}

func PermissionCreate(ctx *fiber.Ctx) error {
	return ctx.Next()
}
