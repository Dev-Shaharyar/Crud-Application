
# CRUD Application

## Project Overview
This backend is built in Go and is structured to follow a clean architecture. Below is a description of the key parts of the project.

---

### `cmd/server`
This directory contains the entry point and server setup:
- **`main.go`**: The main entry point of the application.
- **`server.go`**: Contains the logic for initializing and starting the server.
- **`router.go`**: The router configuration to handle different routes.
- **`todo_routes.go`**: Contains the specific routes for the Todo API.

---

### `config/db`
- **`mongo_init.go`**: Initializes the connection to MongoDB.

---

### `pkg/api/handlers`
This directory contains the handler functions that process API requests:
- **`handlers.go`**: Contains generic handlers for the API.
- **`todo_handlers.go`**: Contains specific handlers for Todo-related operations.

---

### `pkg/application/todo`
This directory contains the business logic for the Todo application:
- **`mapper.go`**: Maps the data between different layers.
- **`todo_data_test.go`**: Contains test data for the Todo application.
- **`todo_service.go`**: The business logic for managing Todos.
- **`todo_service_test.go`**: Contains unit tests for the Todo service.
- **`utils.go`**: Utility functions used across the application.

---

### `pkg/contract/todo`
Defines the contract or API interface for Todo operations:
- **`create_todo.go`**: Defines the structure and logic for creating Todo items.
- **`get_todo.go`**: Defines the structure and logic for retrieving Todo items.
- **`update_todo.go`**: Defines the structure and logic for updating Todo items.

---

### `di`
- **`wire.go`** and **`wire_gen.go`**: Used for dependency injection and wiring dependencies.

---

### `domain/persistence`
Contains repositories for managing data:
- **`mock_todo_repo.go`**: Mock repository for testing.
- **`todo_repo.go`**: The actual repository interface for data persistence.
- **`todo_aggregate`**: Handles the Todo domain logic.
  - **`todo.go`**: Represents the Todo aggregate.
  - **`todo_data.go`**: Represents the Todo data.

---

### `infrastructure/persistence/todo`
Contains the persistence layer for Todo data:
- **`todo_repo.go`**: The repository for interacting with MongoDB.
- **`todo_data_test.go`**: Contains tests for the Todo repository.
- **`todo_models.go`**: Defines the data models for the Todo application.
