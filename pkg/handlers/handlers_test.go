package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Ubivius/microservice-matchmaking/pkg/data"
	"github.com/gorilla/mux"
)

func TestInQueue(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/queue/a2181017-5c53-422b-b6bc-036b27c04fc8", nil)
	response := httptest.NewRecorder()

	queueHandler := NewQueueHandler()

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": "a2181017-5c53-422b-b6bc-036b27c04fc8",
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
	request := httptest.NewRequest(http.MethodGet, "/queue/6e3feed6-96f4-11eb-a8b3-0242ac130003", nil)
	response := httptest.NewRecorder()

	queueHandler := NewQueueHandler()

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": "6e3feed6-96f4-11eb-a8b3-0242ac130003",
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
	request := httptest.NewRequest(http.MethodDelete, "/queue/6e3feed6-96f4-11eb-a8b3-0242ac130003", nil)
	response := httptest.NewRecorder()

	queueHandler := NewQueueHandler()

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": "6e3feed6-96f4-11eb-a8b3-0242ac130003",
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
		UserID: "6e3feed6-96f4-11eb-a8b3-0242ac130003",
	}

	request := httptest.NewRequest(http.MethodPost, "/queue", nil)
	response := httptest.NewRecorder()

	// Add the body to the context since we arent passing through middleware
	ctx := context.WithValue(request.Context(), KeyPlayer{}, body)
	request = request.WithContext(ctx)

	queueHandler := NewQueueHandler()
	queueHandler.AddPlayer(response, request)

	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, but got %d", http.StatusNoContent, response.Code)
	}
}

func TestDeleteExistingPlayer(t *testing.T) {
	request := httptest.NewRequest(http.MethodDelete, "/queue/a2181017-5c53-422b-b6bc-036b27c04fc8", nil)
	response := httptest.NewRecorder()

	queueHandler := NewQueueHandler()

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": "a2181017-5c53-422b-b6bc-036b27c04fc8",
	}
	request = mux.SetURLVars(request, vars)

	queueHandler.Delete(response, request)
	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d but got : %d", http.StatusNoContent, response.Code)
	}
}
