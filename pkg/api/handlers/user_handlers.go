package handlers

import (
	// "github.com/gin-gonic/gin"

	"fmt"
	"net/http"
	"strings"

	uService "github.com/Crud-application/pkg/application/services"
	"github.com/Crud-application/pkg/contracts/user"
	"github.com/gin-gonic/gin"
)

// var _ = eContr.ErrorRes{}

type UserHandler struct {
	userSvc uService.IUserService
}

func NewUserHandler(userService uService.IUserService) *UserHandler {
	return &UserHandler{
		userSvc: userService,
	}
}

// CreateUser creates a new user
func (uh *UserHandler) CreateUser(c *gin.Context) {

	var newUser user.CreateUserReq

	// Parse JSON body
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid input",
			"details": err.Error(), // Include the error message
		})
		return
	}

	fmt.Println("newUser", newUser)
	//update email to lowercase
	newUser.Email = strings.ToLower(newUser.Email)

	user, err := uh.userSvc.CreateUser(c.Request.Context(), &newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// DeleteUser deletes a user by ID
func (uh *UserHandler) DeleteUser(c *gin.Context) {
	// Get the user ID from the URL parameters (assuming it's passed like /user/:userId)
	//fmt.Println("c", c)
	userID := c.Param("userID")
	//fmt.Println("userID", userID)
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	// Pass the user ID to the service layer for deletion
	err := uh.userSvc.DeleteUser(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// GetUser retrieves a user by ID
func (uh *UserHandler) GetUser(c *gin.Context) {
	userID := c.Param("userID")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	user, err := uh.userSvc.GetUser(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No user found with the given ID"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User retrieved successfully",
		"user":    user,
	})
}

// GetUsers retrieves a list of all users
func (uh *UserHandler) GetUsers(c *gin.Context) {
	users, err := uh.userSvc.GetAllUsers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	mapUsers := make([]map[string]interface{}, 0)

	for _, user := range users {
		mapUsers = append(mapUsers, map[string]interface{}{
			"id":          user.ID,
			"name":        user.Name,
			"email":       user.Email,
			"phoneNumber": user.PhoneNumber,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Users retrieved successfully",
		"users":   mapUsers,
	})
}

// PatchUser updates a user's details by ID
func (uh *UserHandler) UpdateUser(c *gin.Context) {
	// Get the user ID from the URL parameters
	userID := c.Param("userID")
	fmt.Printf("userID: %s\n", userID)
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	// Parse the JSON request body into UpdateUserReq
	var req user.UpdateUserReq
	if err := c.BindJSON(&req); err != nil {
		fmt.Printf("Error binding JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invaliddd request body"})
		return
	}

	// Ensure the ID in the request body matches the ID in the URL
	req.ID = userID

	// Call the service layer to update the user
	updatedUser, err := uh.userSvc.UpdateUser(c.Request.Context(), userID, &req)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		}
		return
	}

	// Respond with the updated user details
	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
		"user":    updatedUser,
	})
}
