package main

import (
	"go-task-api/controller"
	"log"

	"github.com/gin-gonic/gin"
)

// main is the entry point of the application.
func main() {
	log.Print("Starting application!")
	serve()
}

// This function:
// 1. Creates a new Gin router instance using `gin.Default()`, which provides built-in middleware for logging and recovery.
// 2. Registers the following routes and their associated handler functions:
//    - `GET /tasks`: Retrieves all tasks by calling the `getTasks` handler.
//    - `GET /tasks/:id`: Retrieves a specific task by ID using the `getTaskByID` handler.
//    - `POST /tasks`: Creates a new task with the `postTasks` handler.
// 3. Starts the HTTP server on `localhost:8080` using `router.Run`.
//
// Notes:
// - The server will listen for incoming HTTP requests on port 8080 of the local machine.
// - The `:id` in the `GET /tasks/:id` route is a dynamic parameter that allows fetching tasks by their unique IDs.
// - With Gin, you can associate a handler with an HTTP method-and-path combination. In this way, you can separately route requests sent to a single path based on the method the client is using.

// Example:
// Once the server is running, you can interact with it using tools like `curl`, Postman, or a browser:
// - `GET http://localhost:8080/tasks`
// - `GET http://localhost:8080/tasks/1`
// - `POST http://localhost:8080/tasks` (with JSON payload)
func serve() {
	router := gin.Default()
	router.GET("/tasks", controller.GetTasks)
	router.GET("/tasks/:id", controller.GetTaskById)
	router.POST("/tasks", controller.PostTasks)
	router.Run("localhost:8080")
}
