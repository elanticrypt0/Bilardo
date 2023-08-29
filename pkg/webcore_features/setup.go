package webcore_features

import (
	"github.com/elanticrypt0/Bilardo/pkg/models"
	"github.com/elanticrypt0/Bilardo/pkg/webcore"
	"github.com/gofiber/fiber/v2"
)

func Setup(c *fiber.Ctx, appcore *webcore.BilardoApp) error {
	migrateModels(appcore)
	return c.SendString("Setup enabled. Models Migrated.")
}

func migrateModels(appcore *webcore.BilardoApp) {
	appcore.Db.AutoMigrate(&models.Category{})
}
