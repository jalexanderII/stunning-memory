package models

import (
	"gorm.io/gorm"
)

type User struct {
	// gorm.Model Embedded Struct, which includes fields ID, CreatedAt, UpdatedAt, DeletedAt
	gorm.Model
	Name        string  `json:"name"`
	Email       *string `json:"email"`
}
