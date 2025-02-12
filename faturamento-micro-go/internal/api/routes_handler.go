package api

import (
	"faturamento-micro-go/internal/entities"
	"faturamento-micro-go/internal/services"

	"github.com/gofiber/fiber/v2"
)

type InvoiceHandler struct {
	service *services.InvoiceService
}

func NewInvoiceHandler(service *services.InvoiceService) *InvoiceHandler {
	return &InvoiceHandler{service: service}
}

func (h *InvoiceHandler) GetAllInvoices(c *fiber.Ctx) error {
	invoices, err := h.service.GetAllInvoices()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(invoices)
}

func (h *InvoiceHandler) CreateInvoice(c *fiber.Ctx) error {
	var invoice entities.Invoice

	if err := c.BodyParser(&invoice); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	createdInvoice, err := h.service.CreateInvoice(invoice)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(createdInvoice)
}

func (h *InvoiceHandler) FindInvoiceByID(c *fiber.Ctx) error {
	id := c.Params("id")

	invoice, err := h.service.FindInvoiceByID(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "invoice not found"})
	}

	return c.Status(fiber.StatusOK).JSON(invoice)
}

func (h *InvoiceHandler) DeleteInvoice(c *fiber.Ctx) error {
	id := c.Params("id")

	deletedProduct, err := h.service.DeleteInvoice(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(deletedProduct)
}
