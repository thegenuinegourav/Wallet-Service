package controller

import (
	"database/sql"
	"encoding/json"
	"github.com/WalletService/cache"
	"github.com/WalletService/service"
	. "github.com/WalletService/model"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type ITransactionController interface {
	GetTransaction(w http.ResponseWriter, r *http.Request)
	GetTransactionsByWalletId(w http.ResponseWriter, r *http.Request)
	GetTransactions(w http.ResponseWriter, r *http.Request)
	GetActiveTransactions(w http.ResponseWriter, r *http.Request)
	PostTransaction(w http.ResponseWriter, r *http.Request)
	PutTransaction(w http.ResponseWriter, r *http.Request)
	UpdateActiveTransactions(w http.ResponseWriter, r *http.Request)
}

type transactionController struct{}

var (
	transactionService service.ITransactionService
	transactionIdempotentCache cache.ITransactionIdempotentCache
)

func NewTransactionController(service service.ITransactionService, idempotent cache.ITransactionIdempotentCache) ITransactionController {
	transactionService = service
	transactionIdempotentCache = idempotent
	return &transactionController{}
}

func (transactionController *transactionController) GetTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid transaction ID")
		return
	}
	transaction, err := transactionService.GetTransactionService(id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Transaction not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, transaction)
}

func (transactionController *transactionController) GetTransactionsByWalletId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid wallet ID")
		return
	}
	transaction, err := transactionService.GetTransactionsByWalletIdService(id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Transactions not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, transaction)
}

func (transactionController *transactionController) GetTransactions(w http.ResponseWriter, r *http.Request) {
	transactions, err := transactionService.GetTransactionsService()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, transactions)
}

func (transactionController *transactionController) GetActiveTransactions(w http.ResponseWriter, r *http.Request) {
	transactions, err := transactionService.GetActiveTransactionsService()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, transactions)
}

func (transactionController *transactionController) PostTransaction(w http.ResponseWriter, r *http.Request) {
	// check for x-idempotency-key
	key, err := transactionIdempotentCache.GetIdempotencyKey(r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	// if this request is not unique, return with idempotent transaction result
	if idempotentTransaction := transactionIdempotentCache.Get(key); idempotentTransaction != nil {
		respondWithJSON(w, http.StatusCreated, idempotentTransaction)
		return
	}
	var transaction Transaction
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&transaction); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid transaction ID")
		return
	}
	defer r.Body.Close()
	res, err := transactionService.PostTransactionService(&transaction, id)
	if err != nil {
		log.Printf("Not able to post Transaction : %s" , err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// Record this request with x-idempotency-key for certain expiry
	transactionIdempotentCache.Set(key, res)
	respondWithJSON(w, http.StatusCreated, res)
}

func (transactionController *transactionController) PutTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid transaction ID")
		return
	}
	var b Transaction
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&b); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	res, err := transactionService.UpdateTransactionService(id, &b)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Transaction not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, res)
}

func (transactionController *transactionController) UpdateActiveTransactions(w http.ResponseWriter, r *http.Request) {
	err := transactionService.UpdateActiveTransactionsService()
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "No Transactions updated!")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"message" : "Transactions marked as inactive."})
}

//func (transactionController *TransactionController) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	id, err := strconv.Atoi(vars["id"])
//	if err != nil {
//		transactionController.RespondWithError(w, http.StatusBadRequest, "Invalid transaction ID")
//		return
//	}
//	err = transactionController.DeleteTransactionService(id)
//	if err != nil {
//		switch err {
//		case sql.ErrNoRows:
//			transactionController.RespondWithError(w, http.StatusNotFound, "Transaction not found")
//		default:
//			transactionController.RespondWithError(w, http.StatusInternalServerError, err.Error())
//		}
//		return
//	}
//	transactionController.RespondWithJSON(w, http.StatusOK, map[string]string{"result" : "success"})
//}
