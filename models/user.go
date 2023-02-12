package models

import (
	"example.com/m/v2/helpers"
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Full_Name string `gorm:"not null" json:"full_name" valid:"required~Full name is required"`
	Email     string `gorm:"not null; unique" json:"email" valid:"required~Email is required, email~Invalid email format"`
	Password  string `gorm:"not null" json:"password" valid:"required~Password is required, stringlength(6|255)~Password has to have a minimum length of 6 characters"`
}

func (u *User) BeforeCreate(g *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
