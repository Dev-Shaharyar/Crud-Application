package user

// @Description GetUserReq is the request structure for get user API call.
type GetUserReq struct {
	ID string `json:"id" example:"tcuZwYseZKNUp8D3tjMkyiZrYGC3" binding:"required"`
} // @name GetUserReq

// @Description GetUserRes is the response structure for get user API call.
type GetUserRes struct {
	ID          string `json:"id" example:"tcuZwYseZKNUp8D3tjMkyiZrYGC3"`
	Name        string `json:"name"`
	Email       string `json:"email" example:"user@example.com"`
	PhoneNumber int64  `json:"phone_number" binding:"-"`
} // @name GetUserRes
