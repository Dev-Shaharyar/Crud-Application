package persistence

import (
	"context"

	useragg "github.com/Crud-application/pkg/domain/userAgg"
	uPersist "github.com/Crud-application/pkg/infrastructure/persistence/user"
)

var _ IUserRepository = (*uPersist.MongoUserRepository)(nil)

// IUserRepository interface
type IUserRepository interface {
	AddUser(ctx context.Context, user *useragg.User) error
	DeleteUser(ctx context.Context, userID string) error
	GetUser(ctx context.Context, userID string) (*useragg.User, error)
	GetAllUser(ctx context.Context) ([]useragg.User, error)
	UpdateUser(ctx context.Context, user *useragg.User) error
}
