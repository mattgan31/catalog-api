package models

type Product struct {
	GormModel
	Product_Name string `gorm:"not_null;unique" json:"product_name" validate:"required"`
	Description  string `gorm:"not_null;unique" json:"description"`
	Image        string `gorm:"not_null;unique" json:"image" validate:"required"`
}
