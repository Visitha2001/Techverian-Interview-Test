package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/visitha2001/backend-go/models"
	"gorm.io/gorm"
)

type ItemHandler struct {
	DB *gorm.DB
}

type Item struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (h *ItemHandler) GetItems(c *fiber.Ctx) error {
	var items []models.Item
	if err := h.DB.Find(&items).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch items",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Items fetched successfully",
		"items":   items,
	})
}

func (h *ItemHandler) GetItem(c *fiber.Ctx) error {
	var item models.Item
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Item ID is required",
		})
	}
	if err := h.DB.First(&item, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch item",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Item fetched successfully",
		"item":    item,
	})
}

func (h *ItemHandler) CreateItem(c *fiber.Ctx) error {
	item := Item{}
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}
	err := h.DB.Create(&item).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create item",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Item created successfully",
		"item":    item,
	})
}

func (h *ItemHandler) DeleteItem(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Item ID is required",
		})
	}

	if err := h.DB.Delete(&models.Item{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete item",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Item deleted successfully",
		"item":    id,
	})
}

func (h *ItemHandler) GetSummary(c *fiber.Ctx) error {
	var items []models.Item
	var totalCost float64
	var averageCost float64

	err := h.DB.Find(&items).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch items for summary",
		})
	}

	for _, item := range items {
		totalCost += item.Price
	}

	averageCost = totalCost / float64(len(items))

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Summary fetched successfully",
		"summary": fiber.Map{
			"totalCost":   totalCost,
			"averageCost": averageCost,
		},
	})
}
