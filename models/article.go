package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title   string `binding:"required"` //必填
	Content string `binding:"required"`
	Preview string `binding:"required"`
	Likes   int    `gorm:"default:0"` //默认值为0
}
