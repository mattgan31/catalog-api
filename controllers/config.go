package controllers

import "gorm.io/gorm"

type InDB struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *InDB {
	return &InDB{
		DB: db,
	}
}
