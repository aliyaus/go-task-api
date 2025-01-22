# Project Overview: Tasks API
An API that allows users to:

* Add tasks.
* Get a list of tasks.
* Get a single task by ID.
* Update a task.
* Delete a task.

# Key ideas and Concepts
* Standard Library: Use net/http for handling HTTP requests and encoding/json for JSON handling.
* Routing: Use a lightweight router like gorilla/mux or chi.
* Persistence: Use an in-memory store (map) for simplicity.
* Idiomatic Go Practices: Proper structuring, error handling, and JSON encoding/decoding.
* Basic unit testing in Golang

# Steps to run project 
1. run `go run main.go`

# Steps to reproduce project
1. Install golang via their website @ https://go.dev/doc/install 
2. `mkdir go-task-api && cd go-task-api`
3. `go mod init go-task-api` to initialize a module for dependency tracking
4. Add contents for `controller/tasks.go`, `model/task.go`, `helper/generate-id.go`
5. run `go get .` to install all the dependencies
6. run application using `go run .`
7. test GET endpoint by running `curl http://localhost:8080/tasks`

# Sample CURLs
GET tasks - `curl http://localhost:8080/tasks`
GET task by ID - `curl http://localhost:8080/tasks/2`
POST task - 
```
curl http://localhost:8080/tasks \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "Test the API","description": "testing the API","completed": true}'
```

# Testing 
1. run `go get -u github.com/stretchr/testify`
2. add `controller/main_test.go` & `controller/tasks_test.go`
3. run `go test -v -cover ./controller` to test


# References
- https://github.com/yemiwebby/gin-gorm-restful-api/tree/main
- https://www.twilio.com/en-us/blog/get-started-testing-api-built-with-golang 