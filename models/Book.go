package models

type Book struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `gorm:"size:255"`
	Author string `gorm:"size:255"`
}
