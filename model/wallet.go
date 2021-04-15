package model

import (
	"github.com/jinzhu/gorm"
)

type Wallet struct {
	gorm.Model
	Balance     float64    	`gorm:"not null;default:0" json:"balance"`
	IsBlock     bool    	`gorm:"not null;default:false" json:"isBlocked"`
	Currency    string    	`gorm:"size:3;not null;" json:"currency"`
	UserID 		uint 		`gorm:"not null;unique;" json:"userId"`
	User		User		`gorm:"foreignKey:UserID"` //This Foreign key tag doesn't work // optional to set user
}
