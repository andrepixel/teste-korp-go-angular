package main

import (
	"faturamento-micro-go/internal/api"
	"faturamento-micro-go/internal/infrastructure"
	"faturamento-micro-go/internal/infrastructure/repositories"
	"faturamento-micro-go/internal/services"
	"fmt"
	"log"
	"os"
	"time"

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

	repository := repositories.NewInvoiceRepository(db)

	service := services.NewInvoiceService(repository)

	api.SetupInvoiceRoutes(app, service)
	
	go simulateFakeError()

	app.Listen(":3000")
}

func fakeError() error {
	return fiber.NewError(fiber.StatusInternalServerError, "Erro fake gerado")
}

func simulateFakeError() {
	for {
		time.Sleep(5 * time.Minute)
		err := fakeError()
		fmt.Println("🚨 Erro simulado:", err)
	}
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
