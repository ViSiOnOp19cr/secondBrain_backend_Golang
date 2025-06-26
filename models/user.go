package models

import "gorm.io/gorm"

type User struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	Email    string    `json:"email" gorm:"unique"`
	Password string    `json:"password"`
	Contents []Content `json:"contents" gorm:"foreignKey:UserID"`
}

type Content struct {
	gorm.Model
	ID     uint   `gorm:"primaryKey" json:"id"`
	UserID uint   `json:"user_id" gorm:"not null"`
	Title  string `json:"title"`
	Type   string `json:"type"`
	Link   string `json:"link"`
	User   User   `json:"-" gorm:"foreignKey:UserID"`
}
