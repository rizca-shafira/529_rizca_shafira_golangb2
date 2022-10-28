package models

import (
	"errors"
	"final_project/mygram/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Email    string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"email,required"`
	Username string `gorm:"not null; unique" json:"username" form:"username" valid:"required"`
	Password string `gorm:"not null" json:"password" form:"password" valid:"required, minstringlength(6)"`
	Age      int    `gorm:"not null" json:"age" form:"age" valid:"required"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	if u.Age < 8 {
		err = errors.New("You are too young!")
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
