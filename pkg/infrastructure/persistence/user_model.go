package user

import (
	uAgg "github.com/Crud-application/pkg/domain/userAgg"
)

// User struct
type User struct {
	ID          string `bson:"_id"`
	Name        string `json:"name" bson:"name"`
	Email       string `json:"email" bson:"email"`
	PhoneNumber int64  `json:"phone_number" bson:"phone_number"`
}

func newUserModel(ua *uAgg.User) *User {
	u := &User{
		ID:          ua.ID,
		Name:        ua.Name,
		Email:       ua.Email,
		PhoneNumber: ua.PhoneNumber,
	}
	return u
}

func (ua *User) toAggregate() (*uAgg.User, error) {
	return &uAgg.User{
		ID:          ua.ID,
		Name:        ua.Name,
		Email:       ua.Email,
		PhoneNumber: ua.PhoneNumber,
	}, nil
}
