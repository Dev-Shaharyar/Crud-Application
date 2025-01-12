package user

import (
	"context"
	"errors"
	"fmt"
	"testing"

	uContr "github.com/Crud-application/pkg/contracts/user"
	uRepoMocks "github.com/Crud-application/pkg/domain/persistence/mocks"
	uAgg "github.com/Crud-application/pkg/domain/userAgg"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock the IUserRepository dependency
	mockRepo := uRepoMocks.NewMockIUserRepository(ctrl)

	// Override the UUID generator for consistency
	generateUUID = func() string {
		return "mocked-uuid"
	}
	defer func() {
		// Reset the UUID generator after tests
		generateUUID = func() string {
			return uuid.New().String()
		}
	}()

	// Initialize the service with the mocked repository
	service := NewUserService(mockRepo)

	// Define test cases
	tests := []struct {
		name          string
		input         *uContr.CreateUserReq
		mockSetup     func()
		expectedRes   *uContr.CreateUserRes
		expectedError error
	}{
		{
			name: "success - user created",
			input: &uContr.CreateUserReq{
				Name:        "John Doe",
				Email:       "johndoe@example.com",
				PhoneNumber: 953285248,
			},
			mockSetup: func() {
				mockRepo.EXPECT().
					AddUser(gomock.Any(), &uAgg.User{
						ID:          "mocked-uuid",
						Name:        "John Doe",
						Email:       "johndoe@example.com",
						PhoneNumber: 953285248,
					}).
					Times(1).
					Return(nil)
			},
			expectedRes: &uContr.CreateUserRes{
				ID:          "mocked-uuid",
				Name:        "John Doe",
				Email:       "johndoe@example.com",
				PhoneNumber: 953285248,
			},
			expectedError: nil,
		},
		{
			name: "failure - repository error",
			input: &uContr.CreateUserReq{
				Name:        "John Doe",
				Email:       "johndoe@example.com",
				PhoneNumber: 953285248,
			},
			mockSetup: func() {
				mockRepo.EXPECT().
					AddUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(errors.New("repository error"))
			},
			expectedRes:   nil,
			expectedError: errors.New("repository error"),
		},
		{
			name: "failure - invalid input",
			input: &uContr.CreateUserReq{
				Name:        "",              // Invalid name
				Email:       "invalid-email", // Invalid email
				PhoneNumber: -1,              // Invalid phone number
			},
			mockSetup:     func() {}, // No repository call expected
			expectedRes:   nil,
			expectedError: errors.New("request validation failed: invalid input"),
		},
	}

	// Run test cases
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Set up the mock expectations
			tc.mockSetup()

			// Call the function under test
			ctx := context.Background()
			res, err := service.CreateUser(ctx, tc.input)

			// Assertions
			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedRes, res)
			}
		})
	}
}

func TestGetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock the IUserRepository dependency
	mockRepo := uRepoMocks.NewMockIUserRepository(ctrl)

	// Initialize the service with the mocked repository
	service := NewUserService(mockRepo)

	// Define test cases
	tests := []struct {
		name          string
		userID        string
		mockSetup     func()
		expectedRes   *uContr.GetUserRes
		expectedError error
	}{
		{
			name:   "success - user found",
			userID: "mocked-uuid",
			mockSetup: func() {
				mockRepo.EXPECT().
					GetUser(gomock.Any(), "mocked-uuid").
					Times(1).
					Return(&uAgg.User{
						ID:          "mocked-uuid",
						Name:        "John Doe",
						Email:       "johndoe@example.com",
						PhoneNumber: 953285248,
					}, nil)
			},
			expectedRes: &uContr.GetUserRes{
				ID:          "mocked-uuid",
				Name:        "John Doe",
				Email:       "johndoe@example.com",
				PhoneNumber: 953285248,
			},
			expectedError: nil,
		},
		{
			name:   "failure - user not found",
			userID: "non-existent-uuid",
			mockSetup: func() {
				mockRepo.EXPECT().
					GetUser(gomock.Any(), "non-existent-uuid").
					Times(1).
					Return(nil, errors.New("user not found"))
			},
			expectedRes:   nil,
			expectedError: errors.New("user not found"),
		},
		{
			name:   "failure - repository error",
			userID: "mocked-uuid",
			mockSetup: func() {
				mockRepo.EXPECT().
					GetUser(gomock.Any(), "mocked-uuid").
					Times(1).
					Return(nil, errors.New("repository error"))
			},
			expectedRes:   nil,
			expectedError: errors.New("repository error"),
		},
	}

	// Run test cases
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Set up the mock expectations
			tc.mockSetup()

			// Call the function under test
			ctx := context.Background()
			res, err := service.GetUser(ctx, tc.userID)
			fmt.Println(res)

			// Assertions
			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedRes, res)
			}
		})
	}
}
func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock the IUserRepository dependency
	mockRepo := uRepoMocks.NewMockIUserRepository(ctrl)

	// Initialize the service with the mocked repository
	service := NewUserService(mockRepo)

	// Define test cases
	tests := []struct {
		name          string
		userID        string
		mockSetup     func()
		expectedError error
	}{
		{
			name:   "success - user deleted",
			userID: "mocked-uuid",
			mockSetup: func() {
				mockRepo.EXPECT().
					DeleteUser(gomock.Any(), "mocked-uuid").
					Times(1).
					Return(nil)
			},
			expectedError: nil,
		},
		{
			name:   "failure - user not found",
			userID: "non-existent-uuid",
			mockSetup: func() {
				mockRepo.EXPECT().
					DeleteUser(gomock.Any(), "non-existent-uuid").
					Times(1).
					Return(errors.New("user not found"))
			},
			expectedError: fmt.Errorf("failed to delete user: %v", errors.New("user not found")),
		},
		{
			name:   "failure - repository error",
			userID: "mocked-uuid",
			mockSetup: func() {
				mockRepo.EXPECT().
					DeleteUser(gomock.Any(), "mocked-uuid").
					Times(1).
					Return(errors.New("repository error"))
			},
			expectedError: fmt.Errorf("failed to delete user: %v", errors.New("repository error")),
		},
	}

	// Run test cases
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Set up the mock expectations
			tc.mockSetup()

			// Call the function under test
			ctx := context.Background()
			err := service.DeleteUser(ctx, tc.userID)

			// Assertions
			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
func TestGetAllUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock the IUserRepository dependency
	mockRepo := uRepoMocks.NewMockIUserRepository(ctrl)

	// Initialize the service with the mocked repository
	service := NewUserService(mockRepo)

	// Define test cases
	tests := []struct {
		name          string
		mockSetup     func()
		expectedRes   []uContr.GetUserRes
		expectedError error
	}{
		{
			name: "success - users found",
			mockSetup: func() {
				mockRepo.EXPECT().
					GetAllUser(gomock.Any()).
					Times(1).
					Return([]uAgg.User{
						{
							ID:          "user-1",
							Name:        "John Doe",
							Email:       "johndoe@example.com",
							PhoneNumber: 953285248,
						},
						{
							ID:          "user-2",
							Name:        "Jane Doe",
							Email:       "janedoe@example.com",
							PhoneNumber: 953285249,
						},
					}, nil)
			},
			expectedRes: []uContr.GetUserRes{
				{
					ID:          "user-1",
					Name:        "John Doe",
					Email:       "johndoe@example.com",
					PhoneNumber: 953285248,
				},
				{
					ID:          "user-2",
					Name:        "Jane Doe",
					Email:       "janedoe@example.com",
					PhoneNumber: 953285249,
				},
			},
			expectedError: nil,
		},
		{
			name: "failure - repository error",
			mockSetup: func() {
				mockRepo.EXPECT().
					GetAllUser(gomock.Any()).
					Times(1).
					Return(nil, errors.New("repository error"))
			},
			expectedRes:   nil,
			expectedError: errors.New("repository error"),
		},
	}

	// Run test cases
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Set up the mock expectations
			tc.mockSetup()

			// Call the function under test
			ctx := context.Background()
			res, err := service.GetAllUsers(ctx)

			// Assertions
			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedRes, res)
			}
		})
	}
}
func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock the IUserRepository dependency
	mockRepo := uRepoMocks.NewMockIUserRepository(ctrl)

	// Initialize the service with the mocked repository
	service := NewUserService(mockRepo)

	// Define test cases
	tests := []struct {
		name          string
		userID        string
		input         *uContr.UpdateUserReq
		mockSetup     func()
		expectedRes   *uContr.UpdateUserRes
		expectedError error
	}{
		{
			name:   "success - user updated",
			userID: "mocked-uuid",
			input: &uContr.UpdateUserReq{
				Name:        stringPtr("John Doe Updated"),
				Email:       stringPtr("johnupdated@example.com"),
				PhoneNumber: int64Ptr(123456789),
			},
			mockSetup: func() {
				mockRepo.EXPECT().
					GetUser(gomock.Any(), "mocked-uuid").
					Times(1).
					Return(&uAgg.User{
						ID:          "mocked-uuid",
						Name:        "John Doe",
						Email:       "johndoe@example.com",
						PhoneNumber: 953285248,
					}, nil)
				mockRepo.EXPECT().
					UpdateUser(gomock.Any(), &uAgg.User{
						ID:          "mocked-uuid",
						Name:        "John Doe Updated",
						Email:       "johnupdated@example.com",
						PhoneNumber: 123456789,
					}).
					Times(1).
					Return(nil)
			},
			expectedRes: &uContr.UpdateUserRes{
				ID:          "mocked-uuid",
				Name:        "John Doe Updated",
				Email:       "johnupdated@example.com",
				PhoneNumber: 123456789,
			},
			expectedError: nil,
		},
		{
			name:   "failure - user not found",
			userID: "non-existent-uuid",
			input: &uContr.UpdateUserReq{
				Name:        stringPtr("John Doe Updated"),
				Email:       stringPtr("johnupdated@example.com"),
				PhoneNumber: int64Ptr(123456789),
			},
			mockSetup: func() {
				mockRepo.EXPECT().
					GetUser(gomock.Any(), "non-existent-uuid").
					Times(1).
					Return(nil, errors.New("user not found"))
			},
			expectedRes:   nil,
			expectedError: fmt.Errorf("user not found: %v", errors.New("user not found")),
		},
		{
			name:   "failure - repository error",
			userID: "mocked-uuid",
			input: &uContr.UpdateUserReq{
				Name:        stringPtr("John Doe Updated"),
				Email:       stringPtr("johnupdated@example.com"),
				PhoneNumber: int64Ptr(123456789),
			},
			mockSetup: func() {
				mockRepo.EXPECT().
					GetUser(gomock.Any(), "mocked-uuid").
					Times(1).
					Return(&uAgg.User{
						ID:          "mocked-uuid",
						Name:        "John Doe",
						Email:       "johndoe@example.com",
						PhoneNumber: 953285248,
					}, nil)
				mockRepo.EXPECT().
					UpdateUser(gomock.Any(), &uAgg.User{
						ID:          "mocked-uuid",
						Name:        "John Doe Updated",
						Email:       "johnupdated@example.com",
						PhoneNumber: 123456789,
					}).
					Times(1).
					Return(errors.New("repository error"))
			},
			expectedRes:   nil,
			expectedError: fmt.Errorf("failed to update user: %v", errors.New("repository error")),
		},
	}

	// Run test cases
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Set up the mock expectations
			tc.mockSetup()

			// Call the function under test
			ctx := context.Background()
			res, err := service.UpdateUser(ctx, tc.userID, tc.input)

			// Assertions
			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedRes, res)
			}
		})
	}
}

func stringPtr(s string) *string {
	return &s
}

func int64Ptr(i int64) *int64 {
	return &i
}
