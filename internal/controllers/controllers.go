package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/utkarshkrsingh/bookStore/internal/config"
	"github.com/utkarshkrsingh/bookStore/internal/models"
)

func CreateBook(res http.ResponseWriter, req *http.Request) {
	log.Printf("Request received at %v...", req.URL)
	if req.Method != http.MethodPost {
		http.Error(res, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var book config.Book
	err := json.NewDecoder(req.Body).Decode(&book)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
	}
	statusCode, err := models.InsertRecord(book)
	if err != nil {
		http.Error(res, fmt.Sprintf("Error: %v", err), statusCode)
		return
	}
}

func GetBooks(res http.ResponseWriter, req *http.Request) {
	log.Printf("Request received at %v...", req.URL)
	if req.Method != http.MethodGet {
		http.Error(res, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}
	books, statusCode, err := models.GetBook()
	if err != nil {
		http.Error(res, err.Error(), statusCode)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(res).Encode(&books); err != nil {
		http.Error(res, "Failed to encode books", http.StatusInternalServerError)
		return
	}
}

func GetBooksByName(res http.ResponseWriter, req *http.Request) {
	log.Printf("Request received at %v...", req.URL)
	if req.Method != http.MethodGet {
		http.Error(res, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var params config.Book
	if err := json.NewDecoder(req.Body).Decode(&params); err != nil {
		http.Error(res, "Failed to decode json", http.StatusInternalServerError)
		return
	}
	bookList, statusCode, err := models.GetBookByName(params.Title, params.Author)
	if err != nil {
		http.Error(res, err.Error(), statusCode)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(res).Encode(&bookList); err != nil {
		http.Error(res, "Failed to encode books", http.StatusInternalServerError)
	}
}

func UpdateByISBN(res http.ResponseWriter, req *http.Request) {
	log.Printf("Request received at %v...", req.URL)
	if req.Method != http.MethodPatch {
		http.Error(res, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var updateData config.Book
	if err := json.NewDecoder(req.Body).Decode(&updateData); err != nil {
		http.Error(res, "Failed to decode json", http.StatusInternalServerError)
		return
	}
	statusCode, err := models.UpdateByISBN(updateData, updateData.ISBN)
	if err != nil {
		http.Error(res, err.Error(), statusCode)
		return
	}
}

func DeleteByISBN(res http.ResponseWriter, req *http.Request) {
	log.Printf("Request received at %v...", req.URL)
	if req.Method != http.MethodDelete {
		http.Error(res, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var bookISBN config.Book
	if err := json.NewDecoder(req.Body).Decode(&bookISBN); err != nil {
		http.Error(res, "Failed to decode json", http.StatusInternalServerError)
		return
	}
	statusCode, err := models.DeleteByISBN(bookISBN.ISBN)
	if err != nil {
		http.Error(res, err.Error(), statusCode)
		return
	}
}
