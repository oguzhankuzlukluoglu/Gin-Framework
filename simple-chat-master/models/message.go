package models

import "github.com/jinzhu/gorm"

// Message model
type Message struct {
	gorm.Model
	Body       string
	ChatID     uint
	SenderID uint
	ReceiverID uint
	IsRead     bool
}
