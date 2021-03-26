package repository

import (
	"github.com/jinzhu/gorm"
	. "github.com/WalletService/model"
)

type IWalletRepository interface{
	GetWalletById(id int) (*Wallet, error)
	GetWalletByUserId(userID int) (*Wallet, error)
	CreateWallet(wallet *Wallet) (*Wallet, error)
	UpdateWallet(wallet *Wallet) (*Wallet, error)
	//DeleteWallet(wallet *Wallet) error
	//GetAllWallets() (*[]Wallet, error)
}

type walletRepository struct {
	DB *gorm.DB
}

func NewWalletRepository(db *gorm.DB) IWalletRepository {
	return &walletRepository{db}
}

func (walletRepository *walletRepository) GetWalletById(id int) (*Wallet, error) {
	var wallet Wallet
	result := walletRepository.DB.Preload("User").First(&wallet,id)
	return &wallet, result.Error
}

func (walletRepository *walletRepository) GetWalletByUserId(userID int) (*Wallet, error) {
	var wallet Wallet
	// use below association approach to avoid preload
	// result := walletRepository.DB.Where("user_id = ?", userID).Find(&wallet)
	// if result.Error != nil {
	// 	return nil, result.Error
	// }
	// err := walletRepository.DB.Model(&wallet).Association("user").Find(&wallet.User).Error
	// return &wallet, err
	result := walletRepository.DB.Where("user_id = ?", userID).Preload("User").First(&wallet)
	return &wallet, result.Error
}

func (walletRepository *walletRepository) CreateWallet(wallet *Wallet) (*Wallet, error) {
	result := walletRepository.DB.Create(wallet)
	return wallet, result.Error
}

func (walletRepository *walletRepository) UpdateWallet(wallet *Wallet) (*Wallet, error) {
	result := walletRepository.DB.Save(wallet)
	return wallet, result.Error
}
//
//func (walletRepository *WalletRepository) DeleteWallet(wallet *Wallet) error {
//	result := walletRepository.DB.Delete(wallet)
//	return result.Error
//}
//
//func (walletRepository *WalletRepository) GetAllWallets() (*[]Wallet, error) {
//	var wallet []Wallet
//	result := walletRepository.DB.Find(&wallet)
//	return &wallet, result.Error
//}