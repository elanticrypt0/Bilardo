package models

import (
	"github.com/elanticrypt0/Bilardo/pkg/webcore"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string `json:"name"`
}

func FindOneCategory(bilardo *webcore.Bilardo, id int) Category {
	var category Category
	bilardo.Db.First(&category, id)
	return category
}

func FindAllCategories(bilardo *webcore.Bilardo) []Category {
	var categories []Category
	// TODO revisar esta parte
	// bilardo.Db.Order("created_at ASC").Find(&categories)
	bilardo.Db.Order("created_at ASC").Find(&categories)
	return categories
}

func CreateCategory(bilardo *webcore.Bilardo, name string) Category {
	category := Category{
		Name: name,
	}
	bilardo.Db.Create(&category)
	return category
}

func UpdateCategory(bilardo *webcore.Bilardo, category Category) Category {
	bilardo.Db.Save(&category)
	return category
}

func DeleteCategory(bilardo *webcore.Bilardo, id int) Category {
	var category Category
	bilardo.Db.First(&category, id)
	bilardo.Db.Delete(&category)
	return category
}
