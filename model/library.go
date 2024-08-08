package model

import (
	"gorm.io/gorm"
	"github.com/google/uuid"
	"github.com/mertburgu/book-management-system/utils"
)

type Library struct {
	gorm.Model
	ID          uuid.UUID `json:"uuid" gorm:"type:uuid;default:uuid_generate_v4();unique"`
	UserID      uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	User        User      `json:"user" gorm:"foreignKey:UserID;references:ID"`
	Name        string    `json:"name" gorm:"not null"`
	Slug        string    `json:"slug" gorm:"unique;not null"`
	IsPublic    bool      `json:"is_public"`
	Description string    `json:"description"`
	Books       []Book    `gorm:"many2many:library_books;"`
}

// BeforeCreate is a GORM hook that gets called before creating a record.
func (l *Library) BeforeCreate(tx *gorm.DB) (err error) {
	if l.ID == uuid.Nil {
		l.ID = uuid.New()
	}
	if l.Name != "" {
		l.Slug = utils.GenerateUniqueSlug(utils.GenerateSlug(l.Name), tx, &Library{})
	}
	return
}
