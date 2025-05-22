package routes

import (
	"github.com/utkarshkrsingh/bookStore/internal/controllers"
	"net/http"
)

var Routers = func(mux *http.ServeMux) {
	mux.HandleFunc("/createBook", controllers.CreateBook)
	mux.HandleFunc("/getBooks", controllers.GetBooks)
	mux.HandleFunc("/getBooksByName", controllers.GetBooksByName)
	mux.HandleFunc("/update", controllers.UpdateByISBN)
	mux.HandleFunc("/delete", controllers.DeleteByISBN)
}
