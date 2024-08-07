package config

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "github.com/mertburgu/book-management-system/model"
)

var DB *gorm.DB

func InitDatabase() {
    var err error
    DB, err = gorm.Open(sqlite.Open("book_management.db"), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database")
    }
    // Auto-migrate models
    DB.AutoMigrate(&model.Book{}, &model.Library{})
}
