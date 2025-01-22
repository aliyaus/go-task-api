package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func router() *gin.Engine {
	router := gin.Default()
	router.GET("/tasks", GetTasks)
	router.GET("/tasks/:id", GetTaskById)
	router.POST("/tasks", PostTasks)

	return router
}

func makeRequest(method, url string, body interface{}) *httptest.ResponseRecorder {
	requestBody, _ := json.Marshal(body)
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(requestBody))

	writer := httptest.NewRecorder()
	router().ServeHTTP(writer, request)
	return writer
}
