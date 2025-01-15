package user

import "github.com/Crud-application/pkg/domain/userAgg"

var TestUserModelData = CreateSampleUser(
	"mocked-uuid",
	"John Doe",
	"johndoe@example.com",
	123456789,
)

func CreateSampleUser(id, name, email string, phone_number int64) userAgg.User {
	return userAgg.User{
		ID:          id,
		Name:        name,
		Email:       email,
		PhoneNumber: phone_number,
	}
}
