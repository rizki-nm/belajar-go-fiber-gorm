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

func GetById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	var user entity.User
	result := database.DB.First(&user, "id = ?", userId)

	if result.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(web.WebResponseFailed{
			Code:    fiber.StatusNotFound,
			Status:  "Not Found",
			Message: "User not found",
		})
	}

	response := web.GetUserResponse{
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
	}

	return ctx.Status(fiber.StatusFound).JSON(web.WebResponseSuccess{
		Code:   fiber.StatusFound,
		Status: "Found",
		Data:   response,
	})

}

func Update(ctx *fiber.Ctx) error {
	userRequest := new(web.UpdateUserEmailRequest)
	err := ctx.BodyParser(userRequest)

	exception.PanicIfNeeded(err)

	errors := validation.Validate(*userRequest)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.WebResponseFailed{
			Code:    fiber.StatusBadRequest,
			Status:  "Bad Request",
			Message: errors,
		})
	}

	userId := ctx.Params("id")

	var user entity.User
	result := database.DB.First(&user, "id = ?", userId)

	if result.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(web.WebResponseFailed{
			Code:    fiber.StatusNotFound,
			Status:  "Not Found",
			Message: "User not found",
		})
	}

	// Update
	user.Email = userRequest.Email

	result = database.DB.Save(&user)

	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.WebResponseFailed{
			Code:    fiber.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: "Failed update email",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponseSuccess{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   user,
	})
}

func Delete(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	var user entity.User
	result := database.DB.First(&user, "id = ?", userId)

	if result.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(web.WebResponseFailed{
			Code:    fiber.StatusNotFound,
			Status:  "Not Found",
			Message: "User not found",
		})
	}

	result = database.DB.Delete(&user)

	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.WebResponseFailed{
			Code:    fiber.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: "Failed delete user",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponseSuccess{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   "User was deleted",
	})
}
