package migration

import (
	"github.com/rizki-nm/belajar-go-fiber-gorm/database"
	"github.com/rizki-nm/belajar-go-fiber-gorm/model/entity"
	"log"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.User{})
	if err != nil {
		log.Println(err)
	}
	log.Println("Database Migrated")
}
