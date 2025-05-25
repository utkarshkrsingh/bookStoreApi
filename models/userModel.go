package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
	UserType string
}

type Book struct {
	gorm.Model
	Title     string  `gorm:"not null" json:"title"`
	Author    string  `gorm:"not null" json:"author"`
	ISBN      string  `gorm:"not null;unique" json:"isbn"`
	Price     float64 `gorm:"type:decimal(10,2);not null" json:"price"`
	Publisher string  `json:"publisher"`
	Published string  `json:"published"`
	Stock     int     `gorm:"default:0" json:"stock"`
}
