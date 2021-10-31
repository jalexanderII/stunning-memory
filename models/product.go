package models

import (
	"gorm.io/gorm"
)

type Product struct {
	// gorm.Model Embedded Struct, which includes fields ID, CreatedAt, UpdatedAt, DeletedAt
	gorm.Model
	Name  string `json:"name"`
	SKU   string `json:"sku"`
	Price uint   `json:"price"`
}
