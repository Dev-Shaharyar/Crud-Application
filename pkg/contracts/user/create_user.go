package user

// @Description CreateUserReq is the request structure for create user API call.
type CreateUserReq struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" example:"user@example.com" binding:"required,email"`
	PhoneNumber int64  `json:"phone_number" binding:"required,numeric"`
} // @name CreateUserReq

// @Description CreateUserRes is the response structure for create user API call.
type CreateUserRes struct {
	ID          string `json:"id" example:"tcuZwYseZKNUp8D3tjMkyiZrYGC3"`
	Name        string `json:"name"`
	Email       string `json:"email" example:"user@example.com"`
	PhoneNumber int64  `json:"phone_number"`
} // @name CreateUserRes
