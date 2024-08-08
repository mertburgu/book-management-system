package utils

import (
    "gorm.io/gorm"
	"strings"
	"unicode"
	"fmt"
)

// GenerateUniqueSlug generates a unique slug for a given baseSlug
func GenerateUniqueSlug(baseSlug string, tx *gorm.DB, model interface{}) string {
	slug := baseSlug
	var count int64
	tx.Model(model).Where("slug LIKE ?", baseSlug+"%").Count(&count)
	if count > 0 {
		slug = fmt.Sprintf("%s-%d", baseSlug, count+1)
	}
	return slug
}

// GenerateSlug generates a slug from a given string
func GenerateSlug(input string) string {
	slug := strings.ToLower(input)
	slug = strings.TrimSpace(slug)
	slug = strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '-' {
			return r
		}
		return -1
	}, slug)
	slug = strings.ReplaceAll(slug, " ", "-")
	return slug
}
