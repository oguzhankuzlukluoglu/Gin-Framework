package models

import "github.com/jinzhu/gorm"

// Chat model
type Chat struct {
	gorm.Model
	Sender     User `gorm:"foreignkey:SenderID"`
	SenderID   uint
	Receiver   User `gorm:"foreignkey:ReceiverID"`
	ReceiverID uint
}
