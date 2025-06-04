package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/utkarshkrsingh/bookStoreApi/controller"
	"github.com/utkarshkrsingh/bookStoreApi/initializers"
	"github.com/utkarshkrsingh/bookStoreApi/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.CoonnectToDB()
	initializers.SyncDatabase()
}

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.POST("/signup", controller.Signup)
	router.POST("/login", controller.Login)
	router.GET("/validate", middleware.RequireAuth, controller.Validate)
	router.GET("/books", controller.GetBooks)
	router.POST("/books", middleware.RequireAuth, controller.InsertBook)
	router.PATCH("/books/:isbn", middleware.RequireAuth, controller.Update)
	router.DELETE("/books/:isbn", middleware.RequireAuth, controller.Delete)

	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
