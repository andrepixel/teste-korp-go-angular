package api

import (
	"controle-micro-go/internal/services"
	"github.com/gofiber/fiber/v2"
)

func SetupProductRoutes(app *fiber.App, productService *services.ProductService) {
	handler := NewProductHandler(productService)

	api := app.Group("/api/products")
	api.Get("/", handler.GetAllProducts)
	api.Post("/", handler.CreateProduct)
	api.Get("/:id", handler.FindProductByID)
	api.Patch("/:id", handler.UpdateProduct)
	api.Delete("/:id", handler.DeleteProduct)
}
