package user

// @Description UpdateUserReq is the request structure for update user API call.
type UpdateUserReq struct {
	ID          string  `json:"id,omitempty" example:"tcuZwYseZKNUp8D3tjMkyiZrYGC3"` // ID of the user (populated later)
	Name        *string `json:"name,omitempty"`                                      // Optional updated name
	Email       *string `json:"email,omitempty" example:"user@example.com"`          // Optional updated email
	PhoneNumber *int64  `json:"phone_number,omitempty"`                              // Optional updated phone number
} // @name UpdateUserReq

// @Description UpdateUserRes is the response structure for update user API call.
type UpdateUserRes struct {
	ID          string `json:"id" example:"tcuZwYseZKNUp8D3tjMkyiZrYGC3"`
	Name        string `json:"name"`
	Email       string `json:"email" example:"user@example.com"`
	PhoneNumber int64  `json:"phone_number"`
} // @name UpdateUserRes
