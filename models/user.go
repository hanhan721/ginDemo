package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique"` //唯一校验标签
	Password string
}
