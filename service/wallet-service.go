package service

import (
	"errors"
	"github.com/WalletService/repository"
	. "github.com/WalletService/model"
)

type IWalletService interface{
	GetWalletService(id int, isUserId bool) (*Wallet, error)
	PostWalletService(wallet *Wallet, userID int) (*Wallet, error)
	UpdateWalletService(updatedWallet *Wallet) (*Wallet, error)
	BlockWalletService(id int) error
	UnBlockWalletService(id int) error
	//DeleteWalletService(id int) error
	//GetWalletsService() (*[]Wallet, error)
}

type walletService struct {}

var (
	walletRepository repository.IWalletRepository
	iUserService IUserService
)

func NewWalletService(repository repository.IWalletRepository, iService IUserService) IWalletService {
	walletRepository = repository
	iUserService = iService
	return &walletService{}
}

func (walletService *walletService) GetWalletService(id int, isUserId bool) (*Wallet, error) {
	if isUserId {
		return walletRepository.GetWalletByUserId(id)
	}
	return walletRepository.GetWalletById(id)
}

func (walletService *walletService) PostWalletService(wallet *Wallet, userID int) (*Wallet, error) {
	user, err := iUserService.GetUserService(userID)
	if err != nil {
		return nil, err
	}
	wallet.UserID = uint(userID)
	wallet.User = *user
	return walletRepository.CreateWallet(wallet)
}

func (walletService *walletService) UpdateWalletService(updatedWallet *Wallet) (*Wallet, error) {
	return walletRepository.UpdateWallet(updatedWallet)
}

func (walletService *walletService) BlockWalletService(id int)  error {
	wallet, err := walletService.GetWalletService(id, false)
	if err != nil {
		return err
	}
	if wallet.IsBlock {
		return errors.New("This wallet is already blocked. Can't block blocked wallet.")
	}
	wallet.IsBlock=true
	_, err = walletRepository.UpdateWallet(wallet)
	return err
}

func (walletService *walletService) UnBlockWalletService(id int)  error {
	wallet, err := walletService.GetWalletService(id, false)
	if err != nil {
		return err
	}
	if !wallet.IsBlock {
		return errors.New("This wallet is already unblocked. Can't unblock unblocked wallet.")
	}
	wallet.IsBlock=false
	_, err = walletRepository.UpdateWallet(wallet)
	return err
}

//func (walletService *WalletService) GetWalletsService() (*[]Wallet, error) {
//	return walletService.GetAllWallets()
//}
//
//func (walletService *WalletService) UpdateWalletService(id int, wallet *Wallet) (*Wallet, error) {
//	res, err := walletService.GetWalletById(id)
//	if err != nil {
//		return nil, err
//	}
//	res.Name = wallet.Name
//	res.Email = wallet.Email
//	res.Mobile = wallet.Mobile
//	return walletService.UpdateWallet(res)
//}
//
//func (walletService *WalletService) DeleteWalletService(id int) error {
//	res, err := walletService.GetWalletById(id)
//	if err != nil {
//		return err
//	}
//	return walletService.DeleteWallet(res)
//}

