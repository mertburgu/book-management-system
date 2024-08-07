package controller

import (
    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"
    "net/http"
    "time"
    "github.com/mertburgu/book-management-system/model"
)

const jwtSecret = "your-secret-key"

// Register handler
func Register(c *gin.Context) {
    // User registration logic
}

// Login handler
func Login(c *gin.Context) {
    var user model.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Validate user and create token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": user.Username,
        "exp":      time.Now().Add(time.Hour * 1).Unix(),
    })

    tokenString, err := token.SignedString([]byte(jwtSecret))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
