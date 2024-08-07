package controller

import (
    "github.com/gin-gonic/gin"
    "github.com/mertburgu/book-management-system/service"
    "github.com/mertburgu/book-management-system/model"
    "net/http"
    "strconv"
)

func CreateBook(c *gin.Context) {
    var book model.Book
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := service.CreateBook(&book); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, book)
}

func GetBooks(c *gin.Context) {
    books, err := service.GetBooks()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, books)
}

func GetBook(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    book, err := service.GetBookByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }
    c.JSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var book model.Book
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    book.ID = uint(id)
    if err := service.UpdateBook(&book); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    if err := service.DeleteBook(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusNoContent, nil)
}
