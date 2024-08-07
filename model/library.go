package model

import (
    "gorm.io/gorm"
)

type Library struct {
    gorm.Model
    UserID uint   `json:"user_id"`
    Name   string `json:"name"`
    IsPublic bool `json:"is_public"`
    Books  []Book `gorm:"many2many:library_books"`
}
