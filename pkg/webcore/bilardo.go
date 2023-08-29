package webcore

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type BilardoApp struct {
	Config AppConfig
	Db     *gorm.DB
	Fiber  *fiber.App
}
