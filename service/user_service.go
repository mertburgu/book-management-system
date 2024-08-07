package service

import (
    "github.com/mertburgu/book-management-system/model"
    "github.com/mertburgu/book-management-system/config"
)

func CreateUser(user *model.User) error {
    return config.DB.Create(user).Error
}

func GetUserByUsername(username string) (model.User, error) {
    var user model.User
    err := config.DB.Where("username = ?", username).First(&user).Error
    return user, err
}
