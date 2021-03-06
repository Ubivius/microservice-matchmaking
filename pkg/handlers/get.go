package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Ubivius/microservice-matchmaking/pkg/data"
)

// InQueue returns a bool that verifies that a player is queued
func (queueHandler *QueueHandler) InQueue(responseWriter http.ResponseWriter, request *http.Request) {
	id := getPlayerID(request)
	log.Info("InQueue request")
	inQueue := data.InQueue(id)
	err := json.NewEncoder(responseWriter).Encode(inQueue)
	if err != nil {
		log.Error(err, "Error serializing inQueue")
		http.Error(responseWriter, "Unable to marshal json", http.StatusInternalServerError)
	}
}
