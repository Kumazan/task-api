# Task API

This is a RESTful Task API application implemented in Go using the Gin web framework. It provides endpoints for managing tasks, including creating, reading, updating, and deleting tasks.

## Features

- RESTful API endpoints for task management
- In-memory data storage
- Dockerized application
- Uses Gin web framework for routing and handling HTTP requests
- Unit tests for repository and service layers

## API Endpoints

- `GET /tasks`: List all tasks
- `POST /tasks`: Create a new task
- `PUT /tasks/:id`: Update an existing task
- `DELETE /tasks/:id`: Delete a task

## Project Structure

```
/task-api
├── cmd
│   └── server
│       └── main.go              # Entry point of the application
├── internal
│   ├── api
│   │   └── handler.go           # HTTP handlers for task endpoints
│   ├── service
│   │   ├── task_service.go      # Business logic for task operations
│   │   └── task_service_test.go # Unit tests for task service
│   ├── repository
│   │   ├── task_repository.go   # In-memory data storage and CRUD operations
│   │   └── task_repository_test.go # Unit tests for task repository
│   └── model
│       └── task.go              # Task data model definition
├── Dockerfile                   # Dockerfile for building the Docker image
├── go.mod                       # Go module file
├── go.sum                       # Go dependencies file
└── README.md                    # Project documentation
```

## Running the Application

1. Build the Docker image:
   ```
   docker build -t task-api .
   ```

2. Run the container:
   ```
   docker run -p 8080:8080 task-api
   ```

The API will be accessible at `http://localhost:8080`.

## Testing

To run the unit tests:

```
go test ./...
```

This will run all tests in the project, including the newly added tests for the repository and service layers.

## Future Improvements

- Add authentication and authorization
- Implement persistent storage (e.g., database)
- Add more comprehensive error handling and logging
- Implement pagination for list endpoints
- Add API documentation (e.g., Swagger)
