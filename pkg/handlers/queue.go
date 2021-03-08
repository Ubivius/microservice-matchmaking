package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// KeyPlayer is a key used for the Player object inside context
type KeyPlayer struct{}

// QueueHandler contains the items common to all queue handler functions
type QueueHandler struct {
	logger *log.Logger
}

// NewQueueHandler returns a pointer to a QueueHandler with the logger passed as a parameter
func NewQueueHandler(logger *log.Logger) *QueueHandler {
	return &QueueHandler{logger}
}

// getPlayerID extracts the player ID from the URL
// The verification of this variable is handled by gorilla/mux
// We panic if it is not valid because that means gorilla is failing
func getPlayerID(request *http.Request) int {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	return id
}
