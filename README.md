# CRUD Application

## Table of Contents

- [Getting Started](#getting-started)
- [Running the Application](#running-the-application)
- [Testing the API](#Testing-the-API)
- [Project Overview](#project-overview)



## Getting Started

Follow these steps to get the project up and running on your local machine.

### Setup

1. **Clone the repository:**

   ```bash
   git clone git@github.com:Dev-Shaharyar/Crud-Application.git
   cd CRUD-2
   

2. **Go to the cmd file:**
   ```bash
   cd cmd



3. **Run below commond:**
```bash
   cd go run main.go



### Running the Application

Once the application is running, you can interact with the API by sending requests to the endpoints. The app runs on the  port `3010`.


## Testing the API

To interact with the CRUD API endpoints, you can use tools like [Postman](https://www.postman.com/)


### API Endpoints

The API supports the following CRUD operations:
    
- **Set the URL = http://localhost:3010/api/

- **Create a User**

  `POST /users`
  
  Request body:
  ```json
  {
    "name":"alam",
    "email":"alam@gmail.com",
    "phone_number":123456789
  }
  ```

  Response body:
  ```json
  {
    "id":"fbhagfbuaygbr",
    "name":"alam",
    "email":"alam@gmail.com",
    "phone_number":123456789
  }
  ```



- **Get all Users**
  
  `GET /users`

  Response:
  ```json
  {
    "message": "Users retrieved successfully",
    "users":[
        {
            "id":"fbhagfbuaygbr",
            "name":"alam",
            "email":"alam@gmail.com",
            "phone_number":123456789
        },
        {
            "id":"ncujbvyewfv",
            "name":"alam bhai",
            "email":"alambhai@gmail.com",
            "phone_number":123456789
        }
    ]
  }
  ```

- **Update a User**
  
  `PATCH/users/{id}`
  
  Request body:
  ```json
  {
    "name": "alam12334",
    "email": "alam@gmail.com",
    "phone_number": 7365725436
  }
  ```

Response body:
  ```json
  {
    "message": "User updated successfully",
    "user": {
        "id": "8db01167-747b-420b-bffb-99cb416f91ac",
        "name": "alam12334",
        "email": "alam@gmail.com",
        "phone_number": 7365725436
    }
  }  
    ```

- **Delete a User**

  `DELETE /users/{id}`

Response body:
  ```json
    {
        "message": "User deleted successfully"
    }
```



## Project Overview
This backend is built in Go and is structured to follow a clean architecture. Below is a description of the key parts of the project.

---

### `cmd`
This directory contains the entry point and server setup:
- **`main.go`**: The main entry point of the application.
- **`server.go`**: Contains the logic for initializing and starting the server.
- **`router.go`**: The router configuration to handle different routes.
- **`user_routes.go`**: Contains the specific routes for the Crud API.

---

### `db`
- **`connection.go`**: Initializes the connection to MongoDB.

---

### `pkg/api/handlers`
This directory contains the handler functions that process API requests:
- **`handlers.go`**: Contains generic handlers for the API.
- **`user_handlers.go`**: Contains specific handlers for user-related operations.

---

### `pkg/application/services`
This directory contains the interfaces of service layer and mocks for that:
- **`user_services`**: Contain Interface of service layer.
- **`mocks/user_services_mock`**: Contains the mocks of service layer interface.


### `pkg/application/user`
This directory contains the business logic for the user application:
- **`mapper.go`**: Maps the data between different layers.
- **`user_data_test.go`**: Contains test data for the Crud application.
- **`user_service.go`**: The business logic for managing User.
- **`user_service_test.go`**: Contains unit tests for the User service.


---

### `pkg/contract/user`
Defines the contract or API interface for user operations:
- **`create_user.go`**: Defines the Request and Response Structure of create user API call.
- **`get_user.go`**: Defines the Request and Response Structure of retrieving user API call.
- **`update_user.go`**: Defines the Request and Response Structure of update user API call.
---

### `di`
- **`wire.go`** and **`wire_gen.go`**: Used for dependency injection and wiring dependencies.

---

### `domain/persistence`
Contains repositories for managing data:
- **`mocks/user_repo_mock.go`**: Mock repository for testing.
- **`user_repo.go`**: The actual repository interface for data persistence.
- **`userAgg`**: Handles the user domain logic.
  - **`user.go`**: Represents the user aggregate.
  - **`user_data.go`**: Represents the user sample data.

---

### `infrastructure/persistence/user`
Contains the persistence layer for user data:
- **`user_repo.go`**: The repository for interacting with MongoDB.
- **`user_model_test.go`**: Contains test data for the user repository.
- **`user_model.go`**: Defines the data models for the user application.
- **`user_repo_test.go`**: Contains unit tests for the User service of Repo layer.
- **`utils`**: contain the utils used in Repo layer.

