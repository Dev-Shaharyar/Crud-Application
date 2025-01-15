package user

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"

	uContr "github.com/Crud-application/pkg/contracts/user"
	mockRepo "github.com/Crud-application/pkg/domain/persistence/mocks"
	"github.com/Crud-application/pkg/domain/userAgg"
	uAgg "github.com/Crud-application/pkg/domain/userAgg"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type fields struct {
	uRepoMocks *mockRepo.MockIUserRepository
}

// Test function for CreateUser
func TestUserService_CreateUser(t *testing.T) {
	req := uContr.CreateUserReq{
		Name:        "John Doe",
		Email:       "johndoe@example.com",
		PhoneNumber: 123456789,
	}
	type args struct {
		ctx context.Context
		req *uContr.CreateUserReq
	}

	type test struct {
		id          int
		name        string
		args        args
		beforeTest  func(f *fields, t *test)
		expectedRes *uContr.CreateUserRes
		wantErr     bool
	}

	tests := []test{
		{
			id:   1,
			name: "CreateUser - success",
			args: args{
				ctx: context.Background(),
				req: &req,
			},
			beforeTest: func(f *fields, t *test) {
				f.uRepoMocks.EXPECT().
					AddUser(gomock.Any(), &uAgg.User{
						ID:          "mocked-uuid",
						Name:        "John Doe",
						Email:       "johndoe@example.com",
						PhoneNumber: 123456789,
					}).
					Return(nil).Times(1)
			},
			expectedRes: &uContr.CreateUserRes{
				ID:          "mocked-uuid",
				Name:        "John Doe",
				Email:       "johndoe@example.com",
				PhoneNumber: 123456789,
			},
			wantErr: false,
		},
		{
			id:   2,
			name: "CreateUser - repository error",
			args: args{
				ctx: context.Background(),
				req: &req,
			},
			beforeTest: func(f *fields, t *test) {
				// mockUser := &uAgg.User{
				// 	ID:          "mocked-uuid",
				// 	Name:        "John Doe",
				// 	Email:       "johndoe@example.com",
				// 	PhoneNumber: 123456789,
				// }

				// Mock AddUser behavior to return error
				f.uRepoMocks.EXPECT().
					AddUser(gomock.Any(), &uAgg.User{
						ID:          "mocked-uuid",
						Name:        "John Doe",
						Email:       "johndoe@example.com",
						PhoneNumber: 123456789,
					}).
					Return(errors.New("repository error")).Times(1)
			},
			expectedRes: nil,
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				uRepoMocks: mockRepo.NewMockIUserRepository(ctrl),
			}

			if tt.beforeTest != nil {
				tt.beforeTest(&f, &tt)
			}

			// Create the UserService instance with the mocked UUID function
			mockUUID := func() string { return "mocked-uuid" } // This is your mocked UUID generator function
			us := NewUserService(f.uRepoMocks, mockUUID)

			// Call the function under test
			got, err := us.CreateUser(tt.args.ctx, tt.args.req)

			// Assertions
			// if tt.expectedError != false {
			// 	assert.Error(t, err)
			// 	assert.Equal(t, tt.expectedError)
			// } else {
			// 	assert.NoError(t, err)
			// 	assert.Equal(t, tt.expectedRes, got)
			// }

			if (err != nil) != tt.wantErr {
				t.Errorf("Test ID %v: Create() error = %v, wantErr %v", tt.id, err, tt.wantErr)
				return
			}
			if tt.expectedRes != nil {
				got.ID = tt.expectedRes.ID
				assert.Equalf(t, tt.expectedRes, got, "Id = %v CreateDevice(%v, %v)", tt.id, tt.args.ctx, tt.args.req)
			}
		})
	}
}
func TestUserService_GetUser(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID string
	}

	type test struct {
		id          int
		name        string
		args        args
		beforeTest  func(f *fields, t *test)
		expectedRes *uContr.GetUserRes
		wantErr     bool
	}

	tests := []test{
		{
			id:   1,
			name: "GetUser - success",
			args: args{
				ctx:    context.Background(),
				userID: "mocked-uuid",
			},
			beforeTest: func(f *fields, t *test) {
				// Mock GetUser behavior
				f.uRepoMocks.EXPECT().
					GetUser(gomock.Any(), gomock.Eq("mocked-uuid")).
					Return(&userAggData1, nil).Times(1)
			},
			expectedRes: &uContr.GetUserRes{
				ID:          "mocked-uuid1",
				Name:        "shaharyar alam",
				Email:       "shaharyar@example.com",
				PhoneNumber: 123456789,
			},
			wantErr: false,
		},
		{
			id:   2,
			name: "GetUser - user not found",
			args: args{
				ctx:    context.Background(),
				userID: "non-existent-id",
			},
			beforeTest: func(f *fields, t *test) {
				// Mock GetUser behavior to return an error (user not found)
				f.uRepoMocks.EXPECT().
					GetUser(gomock.Any(), gomock.Eq("non-existent-id")).
					Return(nil, errors.New("user not found")).Times(1)
			},
			expectedRes: nil,
			wantErr:     true,
		},
		// {
		// 	id:   3,
		// 	name: "GetUser - repository error",
		// 	args: args{
		// 		ctx:    context.Background(),
		// 		userID: "mocked-uuid",
		// 	},
		// 	beforeTest: func(f *fields, t *test) {
		// 		// Mock GetUser behavior to return a repository error
		// 		f.uRepoMocks.EXPECT().
		// 			GetUser(gomock.Any(), gomock.Eq("mocked-uuid")).
		// 			Return(nil, errors.New("repository error")).Times(1)
		// 	},
		// 	expectedRes:   nil,
		// 	expectedError: errors.New("repository error"),
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				uRepoMocks: mockRepo.NewMockIUserRepository(ctrl),
			}

			if tt.beforeTest != nil {
				tt.beforeTest(&f, &tt)
			}

			// Create the UserService instance
			us := NewUserService(f.uRepoMocks, nil)

			// Call the function under test
			got, err := us.GetUser(tt.args.ctx, tt.args.userID)

			// Assertions
			if (err != nil) != tt.wantErr {
				t.Errorf("ID %v GetUserID() error = %v, wantErr %v", tt.id, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.expectedRes) {
				t.Errorf("ID %v GetUserByID() got = %v, want %v", tt.id, got, tt.expectedRes)
			}
		})
	}
}
func TestUserService_DeleteUser(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID string
	}

	type test struct {
		id            int
		name          string
		args          args
		beforeTest    func(f *fields, t *test)
		expectedError error
	}

	tests := []test{
		{
			id:   1,
			name: "DeleteUser - success",
			args: args{
				ctx:    context.Background(),
				userID: "mocked-uuid",
			},
			beforeTest: func(f *fields, t *test) {
				// Mock DeleteUser behavior
				f.uRepoMocks.EXPECT().
					DeleteUser(gomock.Any(), "mocked-uuid").
					Times(1).
					Return(nil)
			},
			expectedError: nil,
		},
		{
			id:   2,
			name: "DeleteUser - user not found",
			args: args{
				ctx:    context.Background(),
				userID: "non-existent-uuid",
			},
			beforeTest: func(f *fields, t *test) {
				// Mock DeleteUser behavior to simulate "user not found" error
				f.uRepoMocks.EXPECT().
					DeleteUser(gomock.Any(), "non-existent-uuid").
					Times(1).
					Return(errors.New("user not found"))
			},
			expectedError: fmt.Errorf("failed to delete user: %v", errors.New("user not found")),
		},
		{
			id:   3,
			name: "DeleteUser - repository error",
			args: args{
				ctx:    context.Background(),
				userID: "mocked-uuid",
			},
			beforeTest: func(f *fields, t *test) {
				// Mock DeleteUser behavior to simulate a repository error
				f.uRepoMocks.EXPECT().
					DeleteUser(gomock.Any(), "mocked-uuid").
					Times(1).
					Return(errors.New("repository error"))
			},
			expectedError: fmt.Errorf("failed to delete user: %v", errors.New("repository error")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Mock the IUserRepository dependency
			f := fields{
				uRepoMocks: mockRepo.NewMockIUserRepository(ctrl),
			}

			// Set up mock expectations
			if tt.beforeTest != nil {
				tt.beforeTest(&f, &tt)
			}

			// Create the UserService instance
			us := NewUserService(f.uRepoMocks, nil)

			// Call the function under test
			err := us.DeleteUser(tt.args.ctx, tt.args.userID)

			// Assertions
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUserService_UpdateUser(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID string
		req    *uContr.UpdateUserReq
	}

	type test struct {
		id            int
		name          string
		args          args
		beforeTest    func(f *fields, t *test)
		expectedRes   *uContr.UpdateUserRes
		expectedError error
	}

	tests := []test{
		{
			id:   1,
			name: "UpdateUser - success",
			args: args{
				ctx:    context.Background(),
				userID: "mocked-uuid",
				req: &uContr.UpdateUserReq{
					Name:        strPtr("Updated Name"),
					Email:       strPtr("updated@example.com"),
					PhoneNumber: intPtr(987654321),
				},
			},
			beforeTest: func(f *fields, t *test) {
				// Mock GetUser behavior to return the existing user
				existingUser := &uAgg.User{
					ID:          "mocked-uuid",
					Name:        "Old Name",
					Email:       "old@example.com",
					PhoneNumber: 123456789,
				}

				// Mock the GetUser and UpdateUser behavior
				f.uRepoMocks.EXPECT().
					GetUser(gomock.Any(), "mocked-uuid").
					Return(existingUser, nil).Times(1)

				updatedUser := &uAgg.User{
					ID:          "mocked-uuid",
					Name:        "Updated Name",
					Email:       "updated@example.com",
					PhoneNumber: 987654321,
				}

				f.uRepoMocks.EXPECT().
					UpdateUser(gomock.Any(), gomock.Eq(updatedUser)).
					Return(nil).Times(1)
			},
			expectedRes: &uContr.UpdateUserRes{
				ID:          "mocked-uuid",
				Name:        "Updated Name",
				Email:       "updated@example.com",
				PhoneNumber: 987654321,
			},
			expectedError: nil,
		},
		{
			id:   2,
			name: "UpdateUser - user not found",
			args: args{
				ctx:    context.Background(),
				userID: "non-existent-uuid",
				req:    &uContr.UpdateUserReq{},
			},
			beforeTest: func(f *fields, t *test) {
				// Mock GetUser behavior to return "user not found"
				f.uRepoMocks.EXPECT().
					GetUser(gomock.Any(), "non-existent-uuid").
					Return(nil, errors.New("user not found"))
			},
			expectedRes:   nil,
			expectedError: fmt.Errorf("user not found: %v", errors.New("user not found")),
		},
		{
			id:   3,
			name: "UpdateUser - repository error on update",
			args: args{
				ctx:    context.Background(),
				userID: "mocked-uuid",
				req: &uContr.UpdateUserReq{
					Name:        strPtr("Updated Name"),
					Email:       strPtr("updated@example.com"),
					PhoneNumber: intPtr(987654321),
				},
			},
			beforeTest: func(f *fields, t *test) {
				// Mock GetUser behavior to return the existing user
				existingUser := &uAgg.User{
					ID:          "mocked-uuid",
					Name:        "Old Name",
					Email:       "old@example.com",
					PhoneNumber: 123456789,
				}

				// Mock the GetUser and UpdateUser behavior
				f.uRepoMocks.EXPECT().
					GetUser(gomock.Any(), "mocked-uuid").
					Return(existingUser, nil).Times(1)

				updatedUser := &uAgg.User{
					ID:          "mocked-uuid",
					Name:        "Updated Name",
					Email:       "updated@example.com",
					PhoneNumber: 987654321,
				}

				// Simulate repository error while updating
				f.uRepoMocks.EXPECT().
					UpdateUser(gomock.Any(), gomock.Eq(updatedUser)).
					Return(errors.New("repository error")).Times(1)
			},
			expectedRes:   nil,
			expectedError: fmt.Errorf("failed to update user: %v", errors.New("repository error")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Mock the IUserRepository dependency
			f := fields{
				uRepoMocks: mockRepo.NewMockIUserRepository(ctrl),
			}

			// Set up mock expectations
			if tt.beforeTest != nil {
				tt.beforeTest(&f, &tt)
			}

			// Create the UserService instance
			us := NewUserService(f.uRepoMocks, nil)

			// Call the function under test
			got, err := us.UpdateUser(tt.args.ctx, tt.args.userID, tt.args.req)

			// Assertions
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedRes, got)
			}
		})
	}
}

func TestUserService_GetAllUsers(t *testing.T) {
	type args struct {
		ctx context.Context
	}

	type test struct {
		id            int
		name          string
		args          args
		beforeTest    func(f *fields, t *test)
		expectedRes   []uContr.GetUserRes
		expectedError error
	}

	tests := []test{
		{
			id:   1,
			name: "GetAllUsers - success",
			args: args{
				ctx: context.Background(),
			},
			beforeTest: func(f *fields, t *test) {
				// Mock GetAllUser behavior to return a list of users
				f.uRepoMocks.EXPECT().
					GetAllUser(gomock.Any()).
					Return([]userAgg.User{
						userAggData1,
						userAggData2,
					}, nil).Times(1)

				// f.uRepoMocks.EXPECT().
				// 	GetAllUser(gomock.Any()).
				// 	Return(&userAggData2, nil).Times(1)
			},

			expectedRes: []uContr.GetUserRes{
				{
					ID:          "mocked-uuid1",
					Name:        "shaharyar alam",
					Email:       "shaharyar@example.com",
					PhoneNumber: 123456789,
				},
				{
					ID:          "mocked-uuid2",
					Name:        "John Doe",
					Email:       "johndoe@example.com",
					PhoneNumber: 987654321,
				},
			},
			expectedError: nil,
		},
		{
			id:   2,
			name: "GetAllUsers - repository error",
			args: args{
				ctx: context.Background(),
			},
			beforeTest: func(f *fields, t *test) {
				// Mock the GetAllUser method to return an error
				f.uRepoMocks.EXPECT().
					GetAllUser(gomock.Any()).
					Return(nil, errors.New("repository error")).Times(1)
			},
			expectedRes:   nil,
			expectedError: errors.New("repository error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Mock the IUserRepository dependency
			f := fields{
				uRepoMocks: mockRepo.NewMockIUserRepository(ctrl),
			}

			// Set up mock expectations
			if tt.beforeTest != nil {
				tt.beforeTest(&f, &tt)
			}

			// Create the UserService instance
			us := NewUserService(f.uRepoMocks, nil)

			// Call the function under test
			got, err := us.GetAllUsers(tt.args.ctx)

			// Assertions
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedRes, got)
			}
		})
	}
}

// Helper functions for pointer dereferencing
func strPtr(s string) *string {
	return &s
}

func intPtr(i int64) *int64 {
	return &i
}
