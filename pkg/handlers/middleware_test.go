package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Ubivius/microservice-matchmaking/pkg/data"
	"github.com/gorilla/mux"
)

func TestValidationMiddlewareWithValidBody(t *testing.T) {
	// Creating request body
	body := &data.Player{
		UserID: 3,
		UserIP: "123.123.123.123",
	}
	bodyBytes, _ := json.Marshal(body)

	request := httptest.NewRequest(http.MethodPost, "/queue", strings.NewReader(string(bodyBytes)))
	response := httptest.NewRecorder()

	queueHandler := NewQueueHandler(NewTestLogger())

	// Create a router for middleware because function attachment is handled by gorilla/mux
	router := mux.NewRouter()
	router.HandleFunc("/queue", queueHandler.AddPlayer)
	router.Use(queueHandler.MiddlewarePlayerValidation)

	// Server http on our router
	router.ServeHTTP(response, request)

	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, but got %d", http.StatusNoContent, response.Code)
	}
}

func TestValidationMiddlewareWithNoUserID(t *testing.T) {
	// Creating request body
	body := &data.Player{
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		t.Error("Body passing to test is not a valid json struct : ", err)
	}

	request := httptest.NewRequest(http.MethodPost, "/queue", strings.NewReader(string(bodyBytes)))
	response := httptest.NewRecorder()

	queueHandler := NewQueueHandler(NewTestLogger())

	// Create a router for middleware because linking is handled by gorilla/mux
	router := mux.NewRouter()
	router.HandleFunc("/queue", queueHandler.AddPlayer)
	router.Use(queueHandler.MiddlewarePlayerValidation)

	// Server http on our router
	router.ServeHTTP(response, request)

	if response.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, response.Code)
	}
	if !strings.Contains(response.Body.String(), "Field validation for 'UserID' failed on the 'required' tag") {
		t.Error("Expected error on field validation for Name but got : ", response.Body.String())
	}
}
