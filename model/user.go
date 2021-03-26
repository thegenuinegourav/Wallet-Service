package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name      string    `gorm:"size:255;not null;" json:"name"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Mobile    string    `gorm:"size:100;not null;" json:"mobile"`
	//Transactions []Transaction.Transaction `gorm:"ForeignKey:UserID"`
}
