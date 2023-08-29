package webcore_features

import (
	"github.com/elanticrypt0/Bilardo/pkg/webcore"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(bilardo *webcore.BilardoApp) {

	// setup
	setup := bilardo.Fiber.Group("/api/setup")
	if bilardo.Config.App_setup_enabled {
		setup.Get("/", func(c *fiber.Ctx) error {
			// return webcore_features.Setup(c)
			return Setup(c, bilardo)
		})
	}

	//status
	status := bilardo.Fiber.Group("/api/status")
	status.Get("/", func(c *fiber.Ctx) error {
		return Status(c)
	})

	status.Get("/getenv", func(c *fiber.Ctx) error {
		return ReadEnv(c)
	})

	seeder := bilardo.Fiber.Group("/api/seeder")
	if bilardo.Config.App_setup_enabled {
		seeder.Get("/api/seed/", func(c *fiber.Ctx) error {
			return Seed(c, bilardo)
		})
		seeder.Get("/api/seed/:table_name", func(c *fiber.Ctx) error {
			return Seed(c, bilardo)
		})
	}
}
