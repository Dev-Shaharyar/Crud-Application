package user

import (
	"context"
	"fmt"
	"log"

	uAgg "github.com/Crud-application/pkg/domain/userAgg"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	client *mongo.Client
}

// NewMongoUserRepository constructor that accepts the MongoDB client
func NewMongoUserRepository(client *mongo.Client) *MongoUserRepository {
	return &MongoUserRepository{
		client: client,
	}
}

// userCollection retrieves the users collection from the database
func (r *MongoUserRepository) userCollection(ctx context.Context) *mongo.Collection {
	// Fetch the collection from the database "crud" and "users"
	return r.client.Database("crud").Collection("users")
}

// AddUser adds a new user to the MongoDB collection
func (r *MongoUserRepository) AddUser(ctx context.Context, user *uAgg.User) error {
	u := toUserModel(user)
	// Insert the document
	result, err := r.userCollection(ctx).InsertOne(ctx, u)
	if err != nil {
		fmt.Printf("Error inserting user: %v", err)
		return err
	}
	fmt.Printf("Inserted user with ID: %s", result.InsertedID)
	return nil
}
func (r *MongoUserRepository) GetUser(ctx context.Context, userID string) (*uAgg.User, error) {
	filter := bson.M{"_id": userID}
	var user *User
	err := r.userCollection(ctx).FindOne(ctx, filter).Decode(&user)
	if err != nil {
		log.Printf("Error getting user: %v", err)
		return nil, err
	}
	return user.toAggregate()
}

func (r *MongoUserRepository) DeleteUser(ctx context.Context, userID string) error {

	filter := bson.M{"_id": userID}

	result, err := r.userCollection(ctx).DeleteOne(ctx, filter)
	if err != nil {
		log.Printf("Error deleting user: %v", err)
		return err
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("no user found with ID: %s", userID)
	}

	log.Printf("Deleted user with ID: %s", userID)
	return nil
}

func (r *MongoUserRepository) GetAllUser(ctx context.Context) ([]uAgg.User, error) {
	// Get all users
	cursor, err := r.userCollection(ctx).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	// Iterate through the cursor
	var users []User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	var res []uAgg.User
	for _, user := range users {
		res = append(res, uAgg.User{
			ID:          user.ID,
			Name:        user.Name,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
		})
	}

	return res, nil
}

func (r *MongoUserRepository) UpdateUser(ctx context.Context, user *uAgg.User) error {
	// Update the user
	filter := bson.M{"_id": user.ID}

	update := bson.M{
		"$set": bson.M{
			"name":        user.Name,
			"email":       user.Email,
			"phoneNumber": user.PhoneNumber,
		},
	}

	_, err := r.userCollection(ctx).UpdateOne(ctx, filter, update)
	if err != nil {
		log.Printf("Error updating user: %v", err)
		return err
	}
	log.Printf("Updated user with ID: %s", user.ID)
	return nil
}
