package controller

import (
	"database/sql"
	"encoding/json"
	. "github.com/WalletService/model"
	"github.com/WalletService/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type IWalletController interface {
	GetWallet(w http.ResponseWriter, r *http.Request, getWithUserId bool)
	PostWallet(w http.ResponseWriter, r *http.Request)
	BlockWallet(w http.ResponseWriter, r *http.Request)
	UnBlockWallet(w http.ResponseWriter, r *http.Request)
}

type walletController struct{}

var (
	walletService service.IWalletService
)

func NewWalletController(service service.IWalletService) IWalletController {
	walletService = service
	return &walletController{}
}

func (walletController *walletController) GetWallet(w http.ResponseWriter, r *http.Request, getWithUserId bool) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}
	wallet, err := walletService.GetWalletService(id, getWithUserId)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Wallet not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, wallet)
}

func (walletController *walletController) PostWallet(w http.ResponseWriter, r *http.Request) {
	var wallet Wallet
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&wallet); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	defer r.Body.Close()
	res, err := walletService.PostWalletService(&wallet, id)
	if err != nil {
		log.Printf("Not able to post Wallet : %s" , err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, res)
}

func (walletController *walletController) BlockWallet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid wallet ID")
		return
	}
	err = walletService.BlockWalletService(id)
	if err != nil {
		log.Printf("Not able to block Wallet : %s" , err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusLocked, map[string]string{"message" : "Wallet is blocked successfully!"})
}

func (walletController *walletController) UnBlockWallet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid wallet ID")
		return
	}
	err = walletService.UnBlockWalletService(id)
	if err != nil {
		log.Printf("Not able to unblock Wallet : %s" , err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusLocked, map[string]string{"message" : "Wallet is unblocked successfully!"})
}



//func (walletController *walletController) PutWallet(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	id, err := strconv.Atoi(vars["id"])
//	if err != nil {
//		respondWithError(w, http.StatusBadRequest, "Invalid wallet ID")
//		return
//	}
//	var b Wallet
//	decoder := json.NewDecoder(r.Body)
//	if err := decoder.Decode(&b); err != nil {
//		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
//		return
//	}
//	defer r.Body.Close()
//	res, err := walletService.UpdateWalletService(id, &b)
//	if err != nil {
//		switch err {
//		case sql.ErrNoRows:
//			respondWithError(w, http.StatusNotFound, "Wallet not found")
//		default:
//			respondWithError(w, http.StatusInternalServerError, err.Error())
//		}
//		return
//	}
//	respondWithJSON(w, http.StatusOK, res)
//}
//
//func (walletController *walletController) DeleteWallet(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	id, err := strconv.Atoi(vars["id"])
//	if err != nil {
//		respondWithError(w, http.StatusBadRequest, "Invalid wallet ID")
//		return
//	}
//	err = walletService.DeleteWalletService(id)
//	if err != nil {
//		switch err {
//		case sql.ErrNoRows:
//			respondWithError(w, http.StatusNotFound, "Wallet not found")
//		default:
//			respondWithError(w, http.StatusInternalServerError, err.Error())
//		}
//		return
//	}
//	respondWithJSON(w, http.StatusOK, map[string]string{"result" : "success"})
//}
//
//func (walletController *walletController) GetWallets(w http.ResponseWriter, r *http.Request) {
//	wallets, err := walletService.GetWalletsService()
//	if err != nil {
//		respondWithError(w, http.StatusInternalServerError, err.Error())
//		return
//	}
//	respondWithJSON(w, http.StatusOK, wallets)
//}
