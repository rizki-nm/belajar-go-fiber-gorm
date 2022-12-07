package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizki-nm/belajar-go-fiber-gorm/database"
	"github.com/rizki-nm/belajar-go-fiber-gorm/exception"
	"github.com/rizki-nm/belajar-go-fiber-gorm/model/entity"
	"github.com/rizki-nm/belajar-go-fiber-gorm/model/web"
	"github.com/rizki-nm/belajar-go-fiber-gorm/validation"
	"log"
)

func GetAll(ctx *fiber.Ctx) error {
	var users []entity.User

	result := database.DB.Find(&users)

	if result.Error != nil {
		log.Println(result.Error)
	}

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponseSuccess{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   users,
	})
}

func Create(ctx *fiber.Ctx) error {
	user := new(web.CreateUserRequest)
	err := ctx.BodyParser(user)

	exception.PanicIfNeeded(err)

	errors := validation.Validate(*user)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.WebResponseFailed{
			Code:    fiber.StatusBadRequest,
			Status:  "Bad Request",
			Message: errors,
		})
	}

	newUser := entity.User{
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
	}

	response := database.DB.Create(&newUser)

	if response.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.WebResponseFailed{
			Code:    fiber.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: "Failed to store data",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(web.WebResponseSuccess{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   newUser,
	})

}
