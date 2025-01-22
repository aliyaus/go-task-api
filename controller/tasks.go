package controller

import (
	"go-task-api/helper"
	"go-task-api/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// tasks slice to seed record task data.
var tasks = []model.Task{
	{ID: "1", Title: "Go shopping", Description: "Go groccery shopping", Completed: true},
	{ID: "2", Title: "Cook dinner", Description: "Cook dinner for family", Completed: false},
	{ID: "3", Title: "Do laundry", Description: "Finish laundry", Completed: false},
}

// getTasks handles HTTP GET requests to retrieve all tasks.
//
// This function:
// 1. Sends a JSON response containing all tasks stored in the global `tasks` slice.
// 2. Uses the `IndentedJSON` method from the Gin framework to format the response as indented JSON for readability.
// 3. Sets the HTTP status code to `200 OK` to indicate a successful response.
//
// Parameters:
// - c (*gin.Context): The Gin context object that provides methods for request and response handling.
//
// Example:
// Assuming the `tasks` slice contains the following tasks:
// [
//
//	{
//	    "title": "Learn Go",
//	    "description": "Complete the Go tutorial",
//	    "completed": false
//	},
//	{
//	    "title": "Build an API",
//	    "description": "Create a REST API in Go",
//	    "completed": true
//	}
//
// ]
// Sending a GET request to this endpoint will respond with:
// [
//
//	{
//	    "title": "Learn Go",
//	    "description": "Complete the Go tutorial",
//	    "completed": false
//	},
//	{
//	    "title": "Build an API",
//	    "description": "Create a REST API in Go",
//	    "completed": true
//	}
//
// ]
//
// Note:
// - If the `tasks` slice is empty, the response will be an empty JSON array `[]` with a `200 OK` status.
func GetTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

// postTasks handles HTTP POST requests to create a new task.
//
// This function:
// 1. Declares a variable `newTask` of type `Task` to hold the incoming task data.
// 2. Parses the JSON payload from the request body into the `newTask` variable using the `BindJSON` method.
//   - If the JSON is invalid or does not match the structure of the `Task` type, the function exits early without adding the task or responding to the client.
//   - The `BindJSON` method automatically validates and decodes the incoming JSON payload into the Go struct.
//
// 3. Appends the parsed `newTask` to the global `tasks` slice, which serves as an in-memory store for tasks.
// 4. Sends a JSON response back to the client with a `201 Created` status code, including the newly created task.
//
// Parameters:
// - c (*gin.Context): The Gin context object that provides methods for request and response handling.
//
// Example:
// Sending a POST request with the following JSON payload:
//
//	{
//	    "title": "Learn Go",
//	    "description": "Complete the Go tutorial",
//	    "completed": false
//	}
//
// Will add the task to the `tasks` slice and respond with:
//
//	{
//	    "title": "Learn Go",
//	    "description": "Complete the Go tutorial",
//	    "completed": false
//	}
func PostTasks(c *gin.Context) {
	var newTask model.Task

	// Call BindJSON to bind the received JSON to newTask.
	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	newId := helper.GenerateID()
	log.Println("Task ID generated =", newId)
	newTask.ID = newId

	// Add the new task to the slice.
	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}

// getTaskByID handles HTTP GET requests to retrieve a single task by its ID.
//
// This function:
// 1. Extracts the `id` parameter from the request URL using `c.Param("id")`.
// 2. Iterates through the global `tasks` slice to find a task with a matching `ID` field.
// 3. If a matching task is found:
//   - Responds with the task data as JSON.
//   - Sets the HTTP status code to `200 OK`.
//
// 4. If no matching task is found:
//   - Responds with a JSON message indicating the task was not found.
//   - Sets the HTTP status code to `404 Not Found`.
//
// Parameters:
//   - c (*gin.Context): The Gin context object, which provides methods for accessing request parameters
//     and sending responses.
//
// Example:
// Assuming the `tasks` slice contains:
// [
//
//	{ "id": "1", "title": "Learn Go", "description": "Complete the Go tutorial", "completed": false },
//	{ "id": "2", "title": "Build an API", "description": "Create a REST API in Go", "completed": true }
//
// ]
//
// Sending a GET request to `/tasks/1` will respond with:
//
//	{
//	    "id": "1",
//	    "title": "Learn Go",
//	    "description": "Complete the Go tutorial",
//	    "completed": false
//	}
//
// Sending a GET request to `/tasks/3` will respond with:
//
//	{
//	    "message": "task not found"
//	}
//
// Note:
// - The `id` parameter is extracted directly from the URL path, and it is case-sensitive.
// - The function assumes the `tasks` slice is populated and accessible globally.
func GetTaskById(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of tasks, looking for
	// an task whose ID value matches the parameter.
	for _, task := range tasks {
		if task.ID == id {
			c.IndentedJSON(http.StatusOK, task)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}
