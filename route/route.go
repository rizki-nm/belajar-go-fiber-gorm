package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizki-nm/belajar-go-fiber-gorm/controller"
	"github.com/rizki-nm/belajar-go-fiber-gorm/middleware"
)

func RouteInit(app *fiber.App) {

	api := app.Group("/api")

	v1 := api.Group("/v1")

	v1.Post("/login", controller.Login)

	v1.Get("/user", middleware.Auth, controller.GetAll)
	v1.Get("/user/:id", controller.GetById)
	v1.Post("/user", controller.Create)
	v1.Put("/user/:id", controller.Update)
	v1.Delete("/user/:id", controller.Delete)
}
