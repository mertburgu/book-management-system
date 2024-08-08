package model

import (
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type Review struct {
	gorm.Model
	ID      uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	BookID  uuid.UUID `json:"book_id" gorm:"type:uuid;not null"`
	UserID  uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	Rating  int       `json:"rating" gorm:"check:rating >= 1 AND rating <= 5"`
	Comment string    `json:"comment"`
	Book    Book      `json:"book" gorm:"foreignKey:BookID"`
	User    User      `json:"user" gorm:"foreignKey:UserID"`
}

// BeforeCreate hook to set UUID
func (r *Review) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}
	return
}
