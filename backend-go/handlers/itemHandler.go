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

func (h *ItemHandler) GetAvarageCost(c *fiber.Ctx) error {
	var items []models.Item
	if err := h.DB.Find(&items).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch items",
		})
	}
	var total float64
	for _, item := range items {
		total += item.Price
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Average cost fetched successfully",
		"average": total / float64(len(items)),
	})
}

func (h *ItemHandler) GetTotalCost(c *fiber.Ctx) error {
	var items []models.Item
	if err := h.DB.Find(&items).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch items",
		})
	}
	var total float64
	for _, item := range items {
		total += item.Price
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Total cost fetched successfully",
		"total":   total,
	})
}
