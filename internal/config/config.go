package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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

var DB *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	db := os.Getenv("DATABASE")
	dbURL := os.Getenv("DATABASE_URL")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("%s:%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbURL, db)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v\n", err)
	}
	// DB.AutoMigrate(&Book{})
}
