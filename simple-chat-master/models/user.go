package models

import "github.com/jinzhu/gorm"

// User model
type User struct {
	gorm.Model
	Name    string `gorm:"type:varchar(20)"`
	Email    string `gorm:"type:varchar(100);unique_index"`
	Password string `json:"-"`
}
