package config

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/mertburgu/book-management-system/model"
)

var DB *gorm.DB

func InitDatabase() {
    var err error
    dsn := "host=db user=user password=password dbname=book_management port=5432 sslmode=disable"
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database")
    }
    // Auto-migrate models
    if err := DB.AutoMigrate(&model.User{}, &model.Book{}, &model.Library{}, &model.Review{}); err != nil {
        panic("Failed to migrate tables: " + err.Error())
    }
}
