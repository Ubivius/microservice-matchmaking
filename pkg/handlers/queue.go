package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// KeyPlayer is a key used for the Player object inside context
type KeyPlayer struct{}

// QueueHandler contains the items common to all queue handler functions
type QueueHandler struct {}

// NewQueueHandler returns a pointer to a QueueHandler with the logger passed as a parameter
func NewQueueHandler() *QueueHandler {
	return &QueueHandler{}
}

// getPlayerID extracts the player ID from the URL
// The verification of this variable is handled by gorilla/mux
func getPlayerID(request *http.Request) string {
	vars := mux.Vars(request)
	id := vars["id"]
	
	return id
}
