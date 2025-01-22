# Project Overview: Tasks API
This is a simple API written in Golang that allows you to do the following
* Add tasks.
* Get a list of tasks.
* Get a single task by ID.

The main purpose of this project is to showcase how to develop a simple backend API service with unit testing in Golang. 

# Key ideas and Concepts
* Rest API in Golang
* Standard Library: Use net/http for handling HTTP requests and encoding/json for JSON handling.
* Idiomatic Go Practices: Proper structuring, error handling, and JSON encoding/decoding.
* Basic unit testing in Golang

# Steps to run project 
* run `go run main.go`

# Steps to reproduce project
1. Install golang via their website @ https://go.dev/doc/install 
2. `mkdir go-task-api && cd go-task-api`
3. `go mod init go-task-api` to initialize a module for dependency tracking
4. Add contents for `controller/tasks.go`, `model/task.go`, `helper/generate-id.go`
5. run `go get .` to install all the dependencies
6. run application using `go run .`
7. test GET endpoint by running `curl http://localhost:8080/tasks`

# Sample CURLs
- GET tasks - `curl http://localhost:8080/tasks`
- GET task by ID - `curl http://localhost:8080/tasks/2`
- POST task - 
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