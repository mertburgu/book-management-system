package api

import (
    "github.com/gin-gonic/gin"
    "github.com/mertburgu/book-management-system/controller"
    "github.com/mertburgu/book-management-system/middleware"
)

func RegisterRoutes(r *gin.Engine) {
    r.POST("/register", controller.Register)
    r.POST("/login", controller.Login)

    authorized := r.Group("/")
    authorized.Use(middleware.AuthMiddleware())
    {
        authorized.GET("/books", controller.GetBooks)
        authorized.POST("/books", controller.CreateBook)
//         authorized.PUT("/books/:id", controller.UpdateBook)
//         authorized.DELETE("/books/:id", controller.DeleteBook)
    }
}
