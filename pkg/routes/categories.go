package routes

import (
	"github.com/elanticrypt0/Bilardo/pkg/features"
	"github.com/elanticrypt0/Bilardo/pkg/webcore"
	"github.com/gofiber/fiber/v2"
)

func categoriesRoutes(bilardo *webcore.Bilardo, api fiber.Router) {
	categories := api.Group("/categories")
	categories.Get("/", func(c *fiber.Ctx) error {
		return features.FindAllCategories(c, bilardo)
	})
	categories.Get("/:id", func(c *fiber.Ctx) error {
		return features.FindOneCategory(c, bilardo)
	})

	categories.Post("/", func(c *fiber.Ctx) error {
		return features.CreateCategory(c, bilardo)
	})

	categories.Put("/:id", func(c *fiber.Ctx) error {
		return features.UpdateCategory(c, bilardo)
	})

	categories.Delete("/:id", func(c *fiber.Ctx) error {
		return features.DeleteCategory(c, bilardo)
	})
}
