package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rizki-nm/belajar-go-fiber-gorm/database"
	"github.com/rizki-nm/belajar-go-fiber-gorm/exception"
	"github.com/rizki-nm/belajar-go-fiber-gorm/model/entity"
	"github.com/rizki-nm/belajar-go-fiber-gorm/model/web"
	"github.com/rizki-nm/belajar-go-fiber-gorm/utils"
	"github.com/rizki-nm/belajar-go-fiber-gorm/validation"
	"time"
)

func Login(ctx *fiber.Ctx) error {
	loginRequest := new(web.LoginRequest)

	err := ctx.BodyParser(loginRequest)

	exception.PanicIfNeeded(err)

	errors := validation.Validate(*loginRequest)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.WebResponseFailed{
			Code:    fiber.StatusBadRequest,
			Status:  "Bad Request",
			Message: errors,
		})
	}

	var user entity.User
	result := database.DB.First(&user, "email = ?", loginRequest.Email)

	if result.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(web.WebResponseFailed{
			Code:    fiber.StatusNotFound,
			Status:  "Not Found",
			Message: "User not found",
		})
	}

	isValid := utils.CheckPasswordHash(loginRequest.Password, user.Password)

	if !isValid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(web.WebResponseFailed{
			Code:    fiber.StatusUnauthorized,
			Status:  "Unauthorized",
			Message: "Password is wrong",
		})
	}

	// Generate JWT
	claims := jwt.MapClaims{}
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["address"] = user.Address
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()

	if user.Email == "admin@test.com" {
		claims["role"] = "admin"
	} else {
		claims["role"] = "user"
	}

	token, err := utils.GenerateToken(&claims)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(web.WebResponseFailed{
			Code:    fiber.StatusUnauthorized,
			Status:  "Unauthorized",
			Message: "Wrong credential",
		})
	}

	return ctx.Status(fiber.StatusFound).JSON(fiber.Map{
		"token": token,
	})
}
