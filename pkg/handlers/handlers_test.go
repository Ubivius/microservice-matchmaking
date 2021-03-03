package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/Ubivius/microservice-matchmaking/pkg/data"
	"github.com/gorilla/mux"
)

// Move to util package in Sprint 9, should be a testing specific logger
func NewTestLogger() *log.Logger {
	return log.New(os.Stdout, "Tests", log.LstdFlags)
}

func TestGetQueue(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/queue", nil)
	response := httptest.NewRecorder()

	queueHandler := NewQueueHandler(NewTestLogger())
	queueHandler.GetQueue(response, request)

	if response.Code != 200 {
		t.Errorf("Expected status code 200 but got : %d", response.Code)
	}
	if !strings.Contains(response.Body.String(), "\"userid\":42") {
		t.Error("Missing elements from expected results")
	}
}

func TestInQueue(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/queue/42", nil)
	response := httptest.NewRecorder()

	queueHandler := NewQueueHandler(NewTestLogger())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": "42",
	}
	request = mux.SetURLVars(request, vars)

	queueHandler.InQueue(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got : %d", http.StatusOK, response.Code)
	}

	var inQueue bool
    err := json.NewDecoder(response.Body).Decode(&inQueue)
	if err != nil {
		t.Error("[ERROR] deserializing player", err)
	} else if !inQueue {
		t.Error("Player is suppose to be in queue")
	}
}

func TestNotInQueue(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/queue/1", nil)
	response := httptest.NewRecorder()

	queueHandler := NewQueueHandler(NewTestLogger())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": "1",
	}
	request = mux.SetURLVars(request, vars)

	queueHandler.InQueue(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got : %d", http.StatusOK, response.Code)
	}

	var inQueue bool
    err := json.NewDecoder(response.Body).Decode(&inQueue)
	if err != nil {
		t.Error("[ERROR] deserializing player", err)
	} else if inQueue {
		t.Error("Player is not suppose to be in queue")
	}
}

func TestDeleteNonExistantPlayer(t *testing.T) {
	request := httptest.NewRequest(http.MethodDelete, "/queue/4", nil)
	response := httptest.NewRecorder()

	queueHandler := NewQueueHandler(NewTestLogger())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": "4",
	}
	request = mux.SetURLVars(request, vars)

	queueHandler.Delete(response, request)
	if response.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d but got : %d", http.StatusNotFound, response.Code)
	}
	if !strings.Contains(response.Body.String(), "Player not found") {
		t.Error("Expected response : Player not found")
	}
}

func TestAddPlayer(t *testing.T) {
	// Creating request body
	body := &data.Player{
		UserID: 1,
	}

	request := httptest.NewRequest(http.MethodPost, "/queue", nil)
	response := httptest.NewRecorder()

	// Add the body to the context since we arent passing through middleware
	ctx := context.WithValue(request.Context(), KeyPlayer{}, body)
	request = request.WithContext(ctx)

	queueHandler := NewQueueHandler(NewTestLogger())
	queueHandler.AddPlayer(response, request)

	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, but got %d", http.StatusNoContent, response.Code)
	}
}

func TestDeleteExistingPlayer(t *testing.T) {
	request := httptest.NewRequest(http.MethodDelete, "/queue/42", nil)
	response := httptest.NewRecorder()

	queueHandler := NewQueueHandler(NewTestLogger())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": "42",
	}
	request = mux.SetURLVars(request, vars)

	queueHandler.Delete(response, request)
	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d but got : %d", http.StatusNoContent, response.Code)
	}
}
