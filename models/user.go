package models

import (
	"gorm.io/gorm"
)

type User struct {
	// gorm.Model Embedded Struct, which includes fields ID, CreatedAt, UpdatedAt, DeletedAt
	gorm.Model
	Name        string  `json:"name" gorm:"index"`
	Email       *string `json:"email" gorm:"index"`
}
