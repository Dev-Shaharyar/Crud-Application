package user

import (
	"context"
	"fmt"

	uCOntr "github.com/Crud-application/pkg/contracts/user"
	uRepo "github.com/Crud-application/pkg/domain/persistence"
)

type UserService struct {
	uRepo        uRepo.IUserRepository
	generateUUID UUIDGenerator
}

func NewUserService(uRepo uRepo.IUserRepository, generateUUID UUIDGenerator) *UserService {
	return &UserService{
		uRepo:        uRepo,
		generateUUID: generateUUID,
	}
}

type UUIDGenerator func() string

func (us *UserService) CreateUser(ctx context.Context, req *uCOntr.CreateUserReq) (*uCOntr.CreateUserRes, error) {
	userID := us.generateUUID()
	//userID := uuid.New().String() //--> getting problem while mocking
	fmt.Printf("hello")
	// Create a new user
	fmt.Println(userID)
	newUser, err := fromCreateUserReq(userID, req)
	//fmt.Printf("hello")
	if err != nil {
		return nil, fmt.Errorf("request validation failed: %v", err)
	}
	err = us.uRepo.AddUser(ctx, newUser)
	if err != nil {
		return nil, err
	}
	return toCreateUserRes(newUser), nil
}

func (us *UserService) DeleteUser(ctx context.Context, userID string) error {
	err := us.uRepo.DeleteUser(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}
	return nil
}

func (us *UserService) GetUser(ctx context.Context, userID string) (*uCOntr.GetUserRes, error) {
	user, err := us.uRepo.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	return toGetUserRes(user), nil
}
func (us *UserService) UpdateUser(ctx context.Context, userID string, req *uCOntr.UpdateUserReq) (*uCOntr.UpdateUserRes, error) {
	// Fetch the existing user from the repository
	existingUser, err := us.uRepo.GetUser(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %v", err)
	}
	// Update fields only if they are provided in the request
	if req.Name != nil {
		existingUser.Name = *req.Name
	}
	if req.Email != nil {
		existingUser.Email = *req.Email
	}
	if req.PhoneNumber != nil {
		existingUser.PhoneNumber = *req.PhoneNumber
	}

	// Save the updated user back to the repository
	err = us.uRepo.UpdateUser(ctx, existingUser)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %v", err)
	}
	return toUpdateUserRes(existingUser), nil
}

func (us *UserService) GetAllUsers(ctx context.Context) ([]uCOntr.GetUserRes, error) {
	// Get users
	users, err := us.uRepo.GetAllUser(ctx)
	if err != nil {
		return nil, err
	}
	return toGetAllUsers(users), nil
}
