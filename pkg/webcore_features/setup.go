package webcore_features

import (
	"github.com/elanticrypt0/Bilardo/pkg/models"
	"github.com/elanticrypt0/Bilardo/pkg/webcore"
	"github.com/gofiber/fiber/v2"
)

func Setup(c *fiber.Ctx, appcore *webcore.Bilardo) error {
	migrateModels(appcore)
	return c.SendString("Setup enabled. Models Migrated.")
}

func migrateModels(appcore *webcore.Bilardo) {
	appcore.Db.AutoMigrate(&models.Category{})
}
