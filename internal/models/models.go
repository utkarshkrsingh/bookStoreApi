package models

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/utkarshkrsingh/bookStore/internal/config"
	"gorm.io/gorm"
)

func InsertRecord(book config.Book) (int, error) {
	var existing config.Book
	err := config.DB.Where("isbn = ?", book.ISBN).First(&existing).Error
	if err == nil {
		return http.StatusConflict, errors.New("ISBN already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("Error checking ISBN uniqueness: %v", err)
		return http.StatusInternalServerError, err
	}

	if err := config.DB.Create(&book).Error; err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return http.StatusConflict, errors.New("ISBN already exists")
		}
		log.Printf("Error creating book: %v", err)
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func GetBook() ([]config.Book, int, error) {
	var books []config.Book
	if err := config.DB.Find(&books).Error; err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return books, http.StatusOK, nil
}

func GetBookByName(title string, author string) ([]config.Book, int, error) {
	var books []config.Book
	if err := config.DB.Where(&config.Book{Title: title, Author: author}).Find(&books).Error; err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return books, http.StatusOK, nil
}

func UpdateByISBN(book config.Book, isbn string) (int, error) {
	var existing config.Book
	if err := config.DB.Where("isbn = ?", isbn).First(&existing).Error; err != nil {
		return http.StatusNotFound, err
	}
	if err := config.DB.Model(&existing).Updates(book).Error; err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func DeleteByISBN(isbn string) (int, error) {
	var book config.Book
	if err := config.DB.Where("isbn = ?", isbn).First(&book).Error; err != nil {
		return http.StatusNotFound, err
	}
	if err := config.DB.Delete(&book).Error; err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
