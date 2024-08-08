package model

import (
	"gorm.io/gorm"
	"github.com/google/uuid"
	"github.com/mertburgu/book-management-system/utils"
)

type Book struct {
	gorm.Model
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Title       string    `json:"title" gorm:"not null"`
	Slug        string    `json:"slug" gorm:"unique;not null"`
	Author      string    `json:"author" gorm:"not null"`
	Publisher   string    `json:"publisher" gorm:"not null"`
	Image       string    `json:"image"`
	Page        int       `json:"page"`
	Year        int       `json:"year"`
	Edition     int       `json:"edition"`
	IsPublic    bool      `json:"is_public"`
	Description string    `json:"description"`
	Language    string    `json:"language"`
}

func (b *Book) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == uuid.Nil {
		b.ID = uuid.New()
	}
	if b.Title != "" {
		b.Slug = utils.GenerateUniqueSlug(utils.GenerateSlug(b.Title), tx, &Book{})
	}
	return
}
