package docs

import "github.com/WalletService/model"

type WalletRequestBody struct {
	Balance     float64    	`gorm:"not null;default:0" json:"balance"`
	Currency    string    	`gorm:"size:3;not null;" json:"currency"`
}

// swagger:route POST /api/v1/user/{ID}/wallet Wallet postWalletParam
// Create wallet for given user id
// responses:
//   200: walletResponse
//   400: error400
//   404: error404
//   500: error500

// Wallet Response
// swagger:response walletResponse
type WalletResponseWrapper struct {
	// in:body
	Body model.Wallet
}

// swagger:parameters postWalletParam
type WalletParamsWrapper struct {
	// Wallet Request Body.
	// in:body
	Body WalletRequestBody
	// User ID
	// in:path
	ID  string  `json:"id"`
}

// swagger:route GET /api/v1/wallet/{ID} Wallet getWalletParam
// Get wallet details for given wallet id
// responses:
//   200: walletResponse
//   400: error400
//   404: error404
//   500: error500

// swagger:parameters getWalletParam
type GetWalletParamsWrapper struct {
	// Wallet ID
	// In: path
	ID string `json:"id"`
}


// swagger:route GET /api/v1/user/{ID}/wallet Wallet getWalletParam2
// Get all wallet details for given user id
// responses:
//   200: getWalletResponse
//   400: error400
//   404: error404
//   500: error500

// swagger:parameters getWalletParam2
type GetWalletParamWrapper struct {
	// User ID
	// In: path
	ID string `json:"id"`
}

// Wallet Response
// swagger:response getWalletResponse
type GetWalletResponseWrapper struct {
	// in:body
	Body []model.Wallet
}

// swagger:route POST /api/v1/wallet/{ID}/block Wallet blockWalletParam
// Block wallet with given id
// responses:
//   200: blockWalletResponse
//   400: error400
//   404: error404
//   500: error500

// Wallet Status Response
// swagger:response blockWalletResponse
type BlockWalletResponse struct {
	// status message
	// in:body
	Body BlockMessage
}

// swagger:parameters blockWalletParam
type BlockWalletParamsWrapper struct {
	// Wallet ID
	// In: path
	ID string `json:"id"`
}

type BlockMessage struct {
	Message string `json:"message"`
}

// swagger:route POST /api/v1/wallet/{ID}/unblock Wallet unblockWalletParam
// Block wallet with given id
// responses:
//   200: blockWalletResponse
//   400: error400
//   404: error404
//   500: error500

// swagger:parameters unblockWalletParam
type UnBlockWalletParamsWrapper struct {
	// Wallet ID
	// In: path
	ID string `json:"id"`
}

