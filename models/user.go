package models

type User struct {
	ID   uint    `gorm:"primaryKey" josn:"id"`
	Email     string `json:"email"`
	Password  string  `json:"password"`
	Content   Content 
}

type Content struct {
	ID         uint          `gorm:"primaryKey" json:"contentid"`
	UserId     uint           `gorm:"unique"`
	Title      string         `josn:"title"`
	Type       string          `json:"type"`
	Link       string          `json:"link"`
}



