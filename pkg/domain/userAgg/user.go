package userAgg

type User struct {
	ID          string
	Name        string
	Email       string
	PhoneNumber int64
}

func NewUser(id, name, email string, phoneNumber int64) (*User, error) {
	return &User{
		ID:          id,
		Name:        name,
		Email:       email,
		PhoneNumber: phoneNumber,
	}, nil
}
