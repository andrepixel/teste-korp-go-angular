package main

import (
	"controle-micro-go/internal/api"
	"controle-micro-go/internal/infrastructure"
	"controle-micro-go/internal/infrastructure/repositories"
	"controle-micro-go/internal/services"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()

	stringConnection := atrributeEnvInString()

	db := infrastructure.ConnectInDatabase(stringConnection)

	if db == nil {
		fmt.Printf("Problem!")
	}

	app := fiber.New()

	repository := repositories.NewProductRepository(db)

	service := services.NewProductService(repository)

	api.SetupProductRoutes(app, service)

	app.Listen(":3000")
}

func loadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading.env file")
	}

	fmt.Println("Variáveis de ambiente carregadas com sucesso!")
}

func atrributeEnvInString() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)
}
