package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"github.com/mertburgu/book-management-system/utils"
)

type User struct {
	gorm.Model
	ID         uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Username   string    `json:"username" gorm:"unique;not null"`
	Slug       string    `json:"slug" gorm:"unique;not null"`
    Password   string    `json:"password" gorm:"not null"`
	Email      string    `json:"email" gorm:"unique;not null"`
	Name       string    `json:"name"`
	Surname    string    `json:"surname"`
	Phone      string    `json:"phone" gorm:"unique;not null"`
	Role       string    `json:"role" gorm:"not null;default:user"`
	Libraries  []Library `json:"libraries" gorm:"foreignKey:UserID"`
	Reviews    []Review  `json:"reviews" gorm:"foreignKey:UserID"`
}

// Role Constants
const (
	RoleUser   = "user"
	RoleEditor = "editor"
	RoleAdmin  = "admin"
)

// BeforeCreate hook to set UUID and Slug
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	if u.Username != "" {
		u.Slug = utils.GenerateUniqueSlug(utils.GenerateSlug(u.Username), tx, &User{})
	}
	return
}
