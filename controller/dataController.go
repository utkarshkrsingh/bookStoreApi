package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/schema"
	"github.com/utkarshkrsingh/bookStoreApi/initializers"
	"github.com/utkarshkrsingh/bookStoreApi/models"
	"gorm.io/gorm"
)

func GetBooks(c *gin.Context) {
	var books []models.Book
	var decoder = schema.NewDecoder()
	var body struct {
		Title string `schema:"title"`
		ISBN  string `schema:"isbn"`
	}

	if err := decoder.Decode(&body, c.Request.URL.Query()); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request: " + err.Error(),
		})

		return
	}

	if body.Title == "" && body.ISBN == "" {
		if err := initializers.DB.Find(&books).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Database error: " + err.Error(),
			})

			return
		}
	} else {
		query := initializers.DB
		if body.Title != "" {
			query = query.Where("title = ?", body.Title)
		}
		if body.ISBN != "" {
			query = query.Where("isbn = ?", body.ISBN)
		}

		err := query.Find(&books).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Database error: " + err.Error(),
			})

			return
		}

		if len(books) == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "No book found with given title and author",
			})

			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data": books,
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
			"error": "Database error: " + err.Error(),
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

	// Prevent ISBN update
	book.ISBN = existing.ISBN

	if err := initializers.DB.Model(&existing).Updates(book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Database error: " + err.Error(),
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
