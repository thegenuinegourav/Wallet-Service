package model

import (
	"github.com/jinzhu/gorm"
)

type Transaction struct {
	gorm.Model
	TxnType   	string 			`gorm:"not null;" json: "txnType"`
	Amount	  	float64			`gorm:"not null;" json:"amount"`
	Active    	bool 			`gorm:"not null;default:true" json:active`
	WalletID 	uint 			`gorm:"not null;" json:"walletId"`
	Wallet	  	Wallet			`gorm:"foreignKey:WalletID;`//This Foreign key tag doesn't work // optional to set user
}
