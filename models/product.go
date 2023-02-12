package models

import "time"

type Product struct {
	ID           uint      `gorm:"primary_key;auto_increment;not_null" json:"id"`
	Product_Name string    `gorm:"not_null;unique" json:"product_name" validate:"required"`
	Description  string    `gorm:"not_null;unique" json:"description"`
	Image        string    `gorm:"not_null;unique" json:"image" validate:"required"`
	Created_At   time.Time `json:"created_at"`
	Updated_At   time.Time `json:"updated_at"`
}
