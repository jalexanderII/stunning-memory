package models

import "gorm.io/gorm"

type Product struct {
	// gorm.Model Embedded Struct, which includes fields ID, CreatedAt, UpdatedAt, DeletedAt
	gorm.Model
	Name  string `gorm:"index" json:"name" `
	SKU   string `gorm:"not null" json:"sku"`
	Price uint   `gorm:"not null" json:"price"`
}
