package user

import "github.com/Crud-application/pkg/domain/userAgg"

var userAggData1 = CreateSampleUser(
	"mocked-uuid1",
	"shaharyar alam",
	"shaharyar@example.com",
	123456789,
)

var userAggData2 = CreateSampleUser(
	"mocked-uuid2",
	"John Doe",
	"johndoe@example.com",
	987654321,
)

func CreateSampleUser(id, name, email string, phone_number int64) userAgg.User {
	return userAgg.User{
		ID:          id,
		Name:        name,
		Email:       email,
		PhoneNumber: phone_number,
	}
}
