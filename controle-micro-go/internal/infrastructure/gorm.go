package infrastructure

import (
	"controle-micro-go/internal/entities"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectInDatabase(stringConnection string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(stringConnection), &gorm.Config{})

	if err != nil {
		log.Fatal("Not is possible connect:", err)
	}

	err = db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error
	if err != nil {
		log.Fatal("Failed to enable uuid-ossp extension:", err)
	}

	err = db.AutoMigrate(&entities.Product{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Printf("Connecting to database")

	return db
}
