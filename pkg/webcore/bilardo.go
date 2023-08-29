package webcore

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Bilardo struct {
	Fiber  *fiber.App
	Db     *gorm.DB
	Config AppConfig
}
