package docs

import "github.com/WalletService/model"

type TransactionRequestBody struct {
	TxnType   	string 			`gorm:"not null;" json: "txnType"`
	Amount	  	float64			`gorm:"not null;" json:"amount"`
}

// swagger:route POST /api/v1/wallet/{ID}/transaction Transaction postTransactionParam
// Create transaction for given wallet id
// responses:
//   200: transactionResponse
//   400: error400
//   404: error404
//   500: error500

// Transaction Response
// swagger:response transactionResponse
type TransactionResponseWrapper struct {
	// in:body
	Body model.Transaction
}

// swagger:parameters postTransactionParam
type TransactionParamsWrapper struct {
	// Transaction Request Body.
	// in:body
	Body TransactionRequestBody
	// Wallet ID
	// in:path
	ID string `json:"id"`
}

// swagger:route GET /api/v1/transaction Transaction getTransactionId
// Fetch all transactions
// responses:
//   200: getTransactionResponse
//   400: error400
//   404: error404
//   500: error500

// Transaction Response
// swagger:response getTransactionResponse
type GetTransactionResponseWrapper struct {
	// in:body
	Body []model.Transaction
}


// swagger:route GET /api/v1/wallet/{ID}/transaction Transaction getTransactionParam
// Fetch transaction associated with given wallet id
// responses:
//   200: getTransactionResponse
//   400: error400
//   404: error404
//   500: error500

// swagger:parameters getTransactionParam
type GetTransactionParamsWrapper struct {
	// Wallet ID
	// In: path
	ID string `json:"id"`
}

// swagger:route GET /api/v1/transaction/active Transaction activeTransactionParam
// Fetch all active transactions
// responses:
//   200: getTransactionResponse
//   400: error400
//   404: error404
//   500: error500

// swagger:route GET /api/v1/transaction/{ID} Transaction getTransactionIDParam
// Fetch transaction for given id
// responses:
//   200: transactionResponse
//   400: error400
//   404: error404
//   500: error500

// swagger:parameters getTransactionIDParam
type GetTransactionParamWrapper struct {
	// Transaction ID
	// In: path
	ID string `json:"id"`
}

// swagger:route PUT /api/v1/transaction/active Transaction putTransactionParam
// Mark all active transaction as inactive
// responses:
//   200: putTransactionResponse
//   400: error400
//   404: error404
//   500: error500

// Transaction Update Response
// swagger:response putTransactionResponse
type DeleteTransactionResponse struct {
	// Update status message
	// in:body
	Body UpdateMessage
}

type UpdateMessage struct {
	Message string `json:"message"`
}

