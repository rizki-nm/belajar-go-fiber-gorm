package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizki-nm/belajar-go-fiber-gorm/database"
	"github.com/rizki-nm/belajar-go-fiber-gorm/database/migration"
	"github.com/rizki-nm/belajar-go-fiber-gorm/route"
)

func main() {
	// Setup Database
	database.DatabaseInit()
	migration.RunMigration()

	app := fiber.New()

	// Setup Routing
	route.RouteInit(app)

	// Start App
	app.Listen(":3000")
}
