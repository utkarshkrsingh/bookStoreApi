package main

import (
	"fmt"
	"github.com/utkarshkrsingh/bookStore/internal/routes"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	routes.Routers(mux)
	fmt.Println("Server starting at port :8000...")
	log.Fatal(http.ListenAndServe(":8000", mux))
}
