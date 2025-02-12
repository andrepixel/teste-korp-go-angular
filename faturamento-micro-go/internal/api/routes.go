package api

import (
	"faturamento-micro-go/internal/services"

	"github.com/gofiber/fiber/v2"
)

func SetupInvoiceRoutes(app *fiber.App, invoiceService *services.InvoiceService) {
	handler := NewInvoiceHandler(invoiceService)

	api := app.Group("/api/invoices")
	api.Get("/", handler.GetAllInvoices)
	api.Post("/", handler.CreateInvoice)
	api.Get("/:id", handler.FindInvoiceByID)
	api.Delete("/:id", handler.DeleteInvoice)
}
