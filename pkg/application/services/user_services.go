// Interfaces

package services

import (
	"context"

	uApp "github.com/Crud-application/pkg/application/user"
	uCOntr "github.com/Crud-application/pkg/contracts/user"
)

// Verify that UserService implements IUserService
var _ IUserService = (*uApp.UserService)(nil)

type IUserService interface {
	//all the user interface methods are defined here
	CreateUser(ctx context.Context, req *uCOntr.CreateUserReq) (*uCOntr.CreateUserRes, error)
	DeleteUser(ctx context.Context, userID string) error
	GetUser(ctx context.Context, userID string) (*uCOntr.GetUserRes, error)
	UpdateUser(ctx context.Context, userID string, req *uCOntr.UpdateUserReq) (*uCOntr.UpdateUserRes, error)
	GetAllUsers(ctx context.Context) ([]uCOntr.GetUserRes, error)
}
