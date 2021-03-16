package models

import (
	"github.com/sawasaki-narumi/calcal-api/utils"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primarykey"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`

	gorm.Model
}

type UserModel struct{}

func (m UserModel) CreateUser(user *User) error {
	return utils.GetDB().Create(user).Error
}

func (m UserModel) FindByEmail(email string) *User {
	user := &User{}
	utils.GetDB().Where("email = ?", email).Find(user)
	return user
}
