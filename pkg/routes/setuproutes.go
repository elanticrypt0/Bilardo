package routes

import (
	"github.com/elanticrypt0/Bilardo/pkg/webcore"
)

func SetupFeaturesRoutes(bilardo *webcore.Bilardo) {
	api := bilardo.Fiber.Group("/api")
	// categories
	categoriesRoutes(bilardo, api)
}
