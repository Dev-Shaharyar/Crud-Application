package server

import (
	"github.com/gin-gonic/gin"
)

// setupUserRoutes registers the basic CRUD routes for user management
func setupUserRoutes(r *gin.RouterGroup, s *HTTPServer) {
	// Create a new user
	r.POST("",
		s.Handlers.UserHandler.CreateUser)

	//Delete a user
	r.DELETE("/:userID",
		s.Handlers.UserHandler.DeleteUser)

	r.GET("", s.Handlers.UserHandler.GetUsers)

	//Get a specific user by ID
	r.GET("/:userID",
		s.Handlers.UserHandler.GetUser)

	//Update a user
	r.PATCH("/:userID",
		s.Handlers.UserHandler.UpdateUser)

}
