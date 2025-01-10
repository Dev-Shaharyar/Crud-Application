package user

import (
	"context"
	"reflect"
	"testing"

	uAgg "github.com/Crud-application/pkg/domain/userAgg"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestMongoUserRepository_AddUser(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	//defer mt.Close()  method for mtest package is removed  not necessary

	type args struct {
		ctx  context.Context
		user *uAgg.User
	}
	tests := []struct {
		id         int
		name       string
		beforeTest func(mt *mtest.T) // Sets up mock responses
		arg        args              // Input user for the AddUser function
		wantErr    bool              // Whether an error is expected
	}{
		{
			id:   1,
			name: "User Added successfully - Success",
			beforeTest: func(mt *mtest.T) {
				mt.AddMockResponses(mtest.CreateSuccessResponse())
			},
			arg: args{
				ctx:  context.Background(),
				user: &uAgg.UserAgg,
			},
			wantErr: false,
		},

		{
			id:   2,
			name: "Create user - MongoDB error - Failure",
			beforeTest: func(mt *mtest.T) {
				mt.AddMockResponses(bson.D{{Key: "ok", Value: 0}})
			},
			arg: args{
				ctx:  context.Background(),
				user: &uAgg.UserAgg,
			},
			wantErr: true,
		},
	}

	// Run test cases
	for _, tt := range tests { // Iterate through test cases
		mt.Run(tt.name, func(mt *mtest.T) {
			// Step 1: Initialize repository with mock MongoDB client
			repo := &MongoUserRepository{
				client: mt.Client,
			}

			// Step 2: Set up mock responses
			if tt.beforeTest != nil {
				tt.beforeTest(mt) // Configure mock behavior
			}

			// Step 3: Call the AddUser method
			err := repo.AddUser(tt.arg.ctx, tt.arg.user)

			// Step 4: Validate error state
			if (err != nil) != tt.wantErr {
				mt.Errorf(
					"Test Case ID: %v | AddUser() error = %v, wantErr = %v",
					tt.id, err, tt.wantErr,
				)
				return // Exit if error state doesn't match
			}

		})
	}
}

func TestMongoUserRepository_GetUser(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	type args struct {
		ctx    context.Context
		userID string
	}

	saExpected, _ := TestUserModelData.toAggregate()
	tests := []struct {
		id         int
		name       string
		beforeTest func(mt *mtest.T)
		arg        args
		want       *uAgg.User
		wantErr    bool
	}{
		{
			id:   1,
			name: "Retrieved User by ID - Success",

			arg: args{
				ctx:    context.Background(),
				userID: TestUserModelData.ID,
			},
			beforeTest: func(mt *mtest.T) {
				mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch,
					bson.D{
						{Key: "_id", Value: TestUserModelData.ID},
						{Key: "name", Value: TestUserModelData.Name},
						{Key: "email", Value: TestUserModelData.Email},
						{Key: "phone_number", Value: TestUserModelData.PhoneNumber},
					}))
			},
			want:    saExpected,
			wantErr: false,
		},
		{
			id:   2,
			name: "User not found by ID - Failure",
			arg: args{
				ctx:    context.Background(),
				userID: "nonexistentID",
			},
			beforeTest: func(mt *mtest.T) {
				// Create a mock response simulating a MongoDB error
				mt.AddMockResponses(bson.D{{Key: "ok", Value: 0}})
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		mt.Run(tt.name, func(mt *mtest.T) {
			repo := &MongoUserRepository{
				client: mt.Client,
			}
			// Step 2: Set up mock responses
			if tt.beforeTest != nil {
				tt.beforeTest(mt) // Configure mock behavior
			}
			// Step 3: Call the AddUser method
			got, err := repo.GetUser(tt.arg.ctx, tt.arg.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ID %v got =%v want = %v", tt.id, got, tt.want)
			}
		})
	}
}

func TestMongoUserRepository_DeleteUser(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	type args struct {
		ctx    context.Context
		userID string
	}
	tests := []struct {
		id         int
		name       string
		args       args
		beforeTest func(mt *mtest.T)
		wantErr    bool
	}{
		{
			id:   1,
			name: "User deleted successfully - Success",
			args: args{
				ctx:    context.Background(),
				userID: TestUserModelData.ID,
			},
			beforeTest: func(mt *mtest.T) {
				mt.AddMockResponses(bson.D{{Key: "ok", Value: 1}, {Key: "acknowledged", Value: true}, {Key: "n", Value: 1}})
			},
			wantErr: false,
		},
		{
			id:   2,
			name: "Delete user - MongoDB error - Failure",
			args: args{
				ctx:    context.Background(),
				userID: "nonexistentUserID",
			},
			beforeTest: func(mt *mtest.T) {
				mt.AddMockResponses(bson.D{{Key: "ok", Value: 0}})
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		mt.Run(tt.name, func(mt *mtest.T) {
			r := &MongoUserRepository{
				client: mt.Client,
			}
			if tt.beforeTest != nil {
				tt.beforeTest(mt) // Configure mock behavior
			}
			if err := r.DeleteUser(tt.args.ctx, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("ID %v DeleteUser() error = %v, wantErr %v", tt.id, err, tt.wantErr)
			}
		})
	}
}
func TestMongoUserRepository_UpdateUser(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	type args struct {
		ctx  context.Context
		user *uAgg.User
	}
	saExpected, _ := TestUserModelData.toAggregate()
	tests := []struct {
		id         int
		name       string
		args       args
		beforeTest func(mt *mtest.T)
		want       *uAgg.User
		wantErr    bool
	}{
		{
			id:   1,
			name: "User updated successfully - Success",
			args: args{
				ctx:  context.Background(),
				user: saExpected,
			},
			beforeTest: func(mt *mtest.T) {
				mt.AddMockResponses(bson.D{
					{Key: "ok", Value: 1},
					{Key: "value", Value: bson.D{
						{Key: "_id", Value: TestUserModelData.ID},
						{Key: "name", Value: TestUserModelData.Name},
						{Key: "email", Value: TestUserModelData.Email},
						{Key: "phone_number", Value: TestUserModelData.PhoneNumber},
					}},
				})
			},
			want:    saExpected,
			wantErr: false,
		},
		{
			id:   2,
			name: "update User Mongo Error - failure",
			beforeTest: func(m *mtest.T) {
				m.AddMockResponses(bson.D{{Key: "ok", Value: 0}})
			},
			args: args{
				ctx:  context.Background(),
				user: saExpected,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		mt.Run(tt.name, func(mt *mtest.T) {
			r := &MongoUserRepository{
				client: mt.Client,
			}
			if tt.beforeTest != nil {
				tt.beforeTest(mt) // Configure mock behavior
			}
			if err := r.UpdateUser(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("ID %v UpdateUser() error = %v, wantErr %v", tt.id, err, tt.wantErr)
			}
		})
	}
}
