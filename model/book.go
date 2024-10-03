package model

type Book struct {
	Id          int    `gorm:"type:int;primary_key" json:"id"`
	Title       string `gorm:"type:varchar(255)" json:"title"`
	Author      string `gorm:"type:varchar(255)" json:"author"`
	Description string `gorm:"type:varchar(512)" json:"description"`
}
