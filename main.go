package main

import (
	"log"

	"github.com/elanticrypt0/Bilardo/pkg/routes"
	"github.com/elanticrypt0/Bilardo/pkg/webcore"
	"github.com/elanticrypt0/Bilardo/pkg/webcore_features"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	app_config := webcore.LoadConfig()

	bilardo := webcore.BilardoApp{
		Fiber:  fiber.New(fiber.Config{}),
		Config: app_config,
		Db:     webcore.Connect2DB(&app_config),
	}

	// CORS
	// necesario para poder utilizar la WUI
	bilardo.Fiber.Use(cors.New())
	bilardo.Fiber.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000,http://localhost:3001,http://localhost:65065",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	bilardo.Fiber.Use(recover.New())

	// features routes
	routes.SetupFeaturesRoutes(&bilardo)
	// webcore features
	webcore_features.SetupRoutes(&bilardo)
	// static routes
	webcore.SetupStaticRoutes(bilardo.Fiber)

	log.Fatal(bilardo.Fiber.Listen(":"+bilardo.Config.App_server_port), "Server is running on port "+bilardo.Config.App_server_port)

}
