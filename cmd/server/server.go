package server

import (
	"log"

	h "github.com/Crud-application/pkg/api/handlers"
	"github.com/Crud-application/pkg/di"
	"github.com/gin-gonic/gin"
)

var (
	BasePath = "/api"
)

type HTTPServer struct {
	Engine   *gin.Engine
	Handlers *h.Handlers
}

// NewServer initializes a new HTTP server with configuration and routes
func NewServer() (*HTTPServer, error) {

	engine := gin.New()
	// Add middlewares (like logger and recovery)
	engine.Use(gin.Logger(), gin.Recovery())

	return &HTTPServer{
		Engine:   engine,
		Handlers: di.InjectHandler(),
	}, nil
}

// Run starts the HTTP server and listens for requests
func (h *HTTPServer) Run() {
	port := ":3010" // Define the port
	log.Printf("Starting server on port %s", port)
	if err := h.Engine.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func (h *HTTPServer) SetupRoutes() {
	SetupPublicRoutes(h)
}
