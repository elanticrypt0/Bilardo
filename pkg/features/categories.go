package features

import (
	"strconv"

	"github.com/elanticrypt0/Bilardo/pkg/models"
	"github.com/elanticrypt0/Bilardo/pkg/webcore"
	"github.com/gofiber/fiber/v2"
)

func FindOneCategory(c *fiber.Ctx, appcore *webcore.Bilardo) error {
	id, _ := strconv.Atoi(c.Params("id", "0"))
	return c.JSON(models.FindOneCategory(appcore, id))
}

func FindAllCategories(c *fiber.Ctx, appcore *webcore.Bilardo) error {
	categories := models.FindAllCategories(appcore)
	return c.JSON(categories)
}

func CreateCategory(c *fiber.Ctx, appcore *webcore.Bilardo) error {
	// name comes from json in body
	cat := new(models.Category)
	c.BodyParser(&cat)
	category := models.CreateCategory(appcore, cat.Name)
	return c.JSON(category)
}

func UpdateCategory(c *fiber.Ctx, appcore *webcore.Bilardo) error {
	id, _ := strconv.Atoi(c.Params("id", "0"))
	category := models.FindOneCategory(appcore, id)
	cat := new(models.Category)
	c.BodyParser(&cat)
	category.Name = cat.Name
	category = models.UpdateCategory(appcore, category)
	return c.JSON(category)
}

func DeleteCategory(c *fiber.Ctx, appcore *webcore.Bilardo) error {
	id, _ := strconv.Atoi(c.Params("id", "0"))
	category := models.DeleteCategory(appcore, id)
	return c.JSON(category)
}
