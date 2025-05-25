package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/utkarshkrsingh/bookStoreApi/initializers"
	"github.com/utkarshkrsingh/bookStoreApi/models"
	"gorm.io/gorm"
)

func GetBooks(c *gin.Context) {
	var books []models.Book
	if err := initializers.DB.Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Database error",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

func GetByName(c *gin.Context) {
	var books []models.Book
	var body struct {
		Title  string
		Author string
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request paylaod: " + err.Error(),
		})

		return
	}

	if body.Title == "" || body.Author == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Both title and author must be provided",
		})
		return
	}

	if err := initializers.DB.Where("title = ? AND author = ?", body.Title, body.Author).Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Database error",
		})

		return
	}

	if len(books) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No book found with given title and author",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"books": books,
	})
}

func GetBookByISBN(c *gin.Context) {
	var book models.Book
	var body struct {
		ISBN string
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request paylaod: " + err.Error(),
		})

		return
	}

	if err := initializers.DB.Where("isbn = ?", body.ISBN).First(&book).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Book not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Database error",
			})
		}

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"book": book,
	})
}

func InsertBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request payload: " + err.Error(),
		})

		return
	}

	var existing models.Book
	if err := initializers.DB.Where("isbn = ?", book.ISBN).First(&existing).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "ISBN already exists",
		})

		return
	} else if err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Database error: ",
		})

		return
	}

	if err := initializers.DB.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to insert book",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Book inserted successfully",
	})
}

func Update(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request payload: " + err.Error(),
		})

		return
	}
	var existing models.Book
	if err := initializers.DB.Where("isbn = ?", book.ISBN).First(&existing).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "ISBN not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Database error: " + err.Error(),
			})
		}
		return
	}

	// Prevent ISBN update
	book.ISBN = existing.ISBN

	if err := initializers.DB.Model(&existing).Updates(book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Database error",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Database updated successfully",
	})
}

func Delete(c *gin.Context) {
	isbn := c.Param("isbn")
	if isbn == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ISBN parameter is required",
		})
		return
	}

	var existing models.Book
	if err := initializers.DB.Where("isbn = ?", isbn).First(&existing).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "ISBN not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Database error: " + err.Error(),
			})
		}
		return
	}

	if err := initializers.DB.Delete(&existing).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Database error: " + err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Book removed successfully",
	})
}
