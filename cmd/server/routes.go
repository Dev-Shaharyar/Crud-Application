package server

// SetupPublicRoutes sets up public routes for user-related resources
func SetupPublicRoutes(h *HTTPServer) {
	crud := h.Engine.Group(BasePath)

	// Define API groups for user-related routes
	userGroup := crud.Group("/users")

	// Set up user-related routes
	setupUserRoutes(userGroup, h)
}
