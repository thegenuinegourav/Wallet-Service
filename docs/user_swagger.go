package docs

import "github.com/WalletService/model"

type UserRequestBody struct {
	Name      string    `gorm:"size:255;not null;" json:"name"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Mobile    string    `gorm:"size:100;not null;" json:"mobile"`
}

// swagger:route POST /api/v1/user User postUserParam
// Create user
// responses:
//   200: userResponse
//   400: error400
//   404: error404
//   500: error500

// User Response
// swagger:response userResponse
type UserResponseWrapper struct {
	// in:body
	Body model.User
}

// swagger:parameters postUserParam
type UserParamsWrapper struct {
	// User Request Body.
	// in:body
	Body UserRequestBody
}

// swagger:route GET /api/v1/user User getUserId
// Fetch all users
// responses:
//   200: getUserResponse
//   400: error400
//   404: error404
//   500: error500

// User Response
// swagger:response getUserResponse
type GetUserResponseWrapper struct {
	// in:body
	Body []model.User
}

// swagger:route PUT /api/v1/user/{ID} User putUserParam
// Update user details for given id
// responses:
//   200: userResponse
//   400: error400
//   404: error404
//   500: error500

// swagger:parameters putUserParam
type PutUserParamsWrapper struct {
	// User Request Body.
	// in:body
	Body UserRequestBody
	// User ID
	// In: path
	ID string `json:"id"`
}

// swagger:route GET /api/v1/user/{ID} User getUserParam
// Get user details for given id
// responses:
//   200: userResponse
//   400: error400
//   404: error404
//   500: error500

// swagger:parameters getUserParam
type GetUserParamsWrapper struct {
	// User ID
	// In: path
	ID string `json:"id"`
}

// swagger:route DELETE /api/v1/user/{ID} User delUserParam
// Delete user with given id
// responses:
//   200: delUserResponse
//   400: error400
//   404: error404
//   500: error500

// User delete Response
// swagger:response delUserResponse
type DeleteUserResponse struct {
	// Delete status message
	// in:body
	Body DeleteMessage
}

// swagger:parameters delUserParam
type DelUserParamsWrapper struct {
	// User ID
	// In: path
	ID string `json:"id"`
}

type DeleteMessage struct {
	Message string `json:"message"`
}

