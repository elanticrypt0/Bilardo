package webcore

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type BilardoApp struct {
	Fiber  *fiber.App
	Db     *gorm.DB
	Config AppConfig
}
