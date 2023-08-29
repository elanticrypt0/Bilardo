package models

import (
	"github.com/elanticrypt0/Bilardo/pkg/webcore"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string `json:"name"`
}

func FindOneCategory(bilardo *webcore.BilardoApp, id int) Category {
	var category Category
	bilardo.Db.First(&category, id)
	return category
}

func FindAllCategories(bilardo *webcore.BilardoApp) []Category {
	var categories []Category
	// TODO revisar esta parte
	// bilardo.Db.Order("created_at ASC").Find(&categories)
	bilardo.Db.Order("created_at ASC").Find(&categories)
	return categories
}

func CreateCategory(bilardo *webcore.BilardoApp, name string) Category {
	category := Category{
		Name: name,
	}
	bilardo.Db.Create(&category)
	return category
}

func UpdateCategory(bilardo *webcore.BilardoApp, category Category) Category {
	bilardo.Db.Save(&category)
	return category
}

func DeleteCategory(bilardo *webcore.BilardoApp, id int) Category {
	var category Category
	bilardo.Db.First(&category, id)
	bilardo.Db.Delete(&category)
	return category
}
