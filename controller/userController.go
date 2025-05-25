package controller

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/utkarshkrsingh/bookStoreApi/initializers"
	"github.com/utkarshkrsingh/bookStoreApi/models"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	// Get the user email/password of the request body
	var body struct {
		Email    string
		Password string
		UserType string
	}

	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	if body.Email == "" || body.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email and password are required",
		})

		return
	}

	if body.UserType != "customer" {
		body.UserType = "customer"
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

		return
	}

	// Create the user
	user := models.User{Email: body.Email, Password: string(hash), UserType: body.UserType}
	if err := initializers.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "Signup successful",
	})

}

func Login(c *gin.Context) {
	// Get the email and password from the req body
	var body struct {
		Email    string
		Password string
	}

	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}
	// Look up requested user
	var user models.User
	initializers.DB.Where("email = ?", body.Email).First(&user)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})

		return
	}

	// Compare sent in password with saved user password hash
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})

		return
	}

	// Generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":      user.ID,
		"userType": user.UserType,
		"exp":      time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	secret := os.Getenv("SECRET")
	if secret == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Token signing secret is not configured",
		})

		return
	}

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})

		return
	}

	// send it back
	c.SetSameSite(http.SameSiteLaxMode)
	secure := gin.Mode() == gin.ReleaseMode
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", secure, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"email":   user.Email,
	})
}

func Validate(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not exists",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User is authenticated",
		"user":    user,
	})
}
