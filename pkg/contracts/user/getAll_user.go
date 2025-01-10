package user

// use query param for limit and offset
// @Description GetAllUsersReq is the request structure for getting all users API call.
type GetAllUsersReq struct {
	Limit  int `json:"limit" example:"10" binding:"omitempty"` // Optional limit for number of users to fetch
	Offset int `json:"offset" example:"0" binding:"omitempty"` // Optional offset for pagination
} // @name GetAllUsersReq

// @Description GetAllUsersRes is the response structure for getting all users API call.
type GetAllUsersRes struct {
	Users []User `json:"users"`               // List of users
	Total int    `json:"total" example:"100"` // Total count of users in the system
} // @name GetAllUsersRes

// @Description User is the structure representing a user in the GetAllUsers response.
type User struct {
	ID          string `json:"id" example:"tcuZwYseZKNUp8D3tjMkyiZrYGC3"`
	Name        string `json:"name"`
	Email       string `json:"email" example:"user@example.com"`
	PhoneNumber int64  `json:"phone_number" binding:"-"`
} // @name User
