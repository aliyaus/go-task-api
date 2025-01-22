package controller

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTasks(t *testing.T) {
	writer := makeRequest("GET", "/tasks", nil)
	assert.Equal(t, http.StatusOK, writer.Code)
	assert.Contains(t, writer.Body.String(), `"title": "Go shopping"`)
}

func TestGetTaskById(t *testing.T) {
	writer := makeRequest("GET", "/tasks/1", nil)
	assert.Equal(t, http.StatusOK, writer.Code)
	assert.Contains(t, writer.Body.String(), `"title": "Go shopping"`)
}

func TestPostTask(t *testing.T) {
	newTask := map[string]interface{}{
		"title":       "New Task",
		"description": "This is a new task",
		"completed":   false,
	}
	writer := makeRequest("POST", "/tasks", newTask)
	assert.Equal(t, http.StatusCreated, writer.Code)
}
