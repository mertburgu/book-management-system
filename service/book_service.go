package service

import (
    "github.com/mertburgu/book-management-system/config"
    "github.com/mertburgu/book-management-system/model"
)

func CreateBook(book *model.Book) error {
    return config.DB.Create(book).Error
}

func GetBooks() ([]model.Book, error) {
    var books []model.Book
    err := config.DB.Find(&books).Error
    return books, err
}

/* func GetBookByID(id uint) (model.Book, error) {
    var book model.Book
    err := config.DB.First(&book, id).Error
    return book, err
} */

func UpdateBook(book *model.Book) error {
    return config.DB.Save(book).Error
}

/* func DeleteBook(id uint) error {
    return config.DB.Delete(&model.Book{}, id).Error
} */
