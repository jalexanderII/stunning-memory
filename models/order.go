package models

import "gorm.io/gorm"

type Order struct {
	// gorm.Model Embedded Struct, which includes fields ID, CreatedAt, UpdatedAt, DeletedAt
	gorm.Model
	// A belongs to association sets up a one-to-one connection with another model,
	// such that each instance of the declaring model “belongs to” one instance of the other model.
	ProductRef int `json:"product_id"`
	Product Product `gorm:"foreignKey:ProductRef"`
	UserRef int `json:"user_id"`
	User User `gorm:"foreignKey:UserRef"`
}