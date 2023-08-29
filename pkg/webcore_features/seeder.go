package webcore_features

import (
	"encoding/json"
	"log"

	"github.com/elanticrypt0/Bilardo/pkg/models"
	"github.com/elanticrypt0/Bilardo/pkg/webcore"
	"github.com/elanticrypt0/Bilardo/pkg/webcore/helpers"
	"github.com/gofiber/fiber/v2"
)

func Seed(c *fiber.Ctx, bilardo *webcore.BilardoApp) error {
	// seedCategories()
	return c.JSON("OK")
}

func seedTable(table_name string) {

}

func seedCategories(bilardo *webcore.BilardoApp) {
	file := helpers.ReadJsonFile("categories")
	cat := []models.Category{}
	err := json.Unmarshal(file, &cat)
	if err != nil {
		log.Fatal("Cant unmarshal json", err)
	}
	for _, category := range cat {
		models.CreateCategory(bilardo, category.Name)
	}
	log.Println("Categories seeded")
}
