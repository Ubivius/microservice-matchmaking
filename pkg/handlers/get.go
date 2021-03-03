package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Ubivius/microservice-matchmaking/pkg/data"
)

// GetQueue returns the full list of players in queue
func (queueHandler *QueueHandler) GetQueue(responseWriter http.ResponseWriter, request *http.Request) {
	queueHandler.logger.Println("Handle GET queue")
	queue := data.GetQueue()
	err := json.NewEncoder(responseWriter).Encode(queue)
	if err != nil {
		queueHandler.logger.Println("[ERROR] serializing queue", err)
		http.Error(responseWriter, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// InQueue returns a bool that verifies that a player is queued
func (queueHandler *QueueHandler) InQueue(responseWriter http.ResponseWriter, request *http.Request) {
	id := getPlayerID(request)
	queueHandler.logger.Println("Handle DELETE player", id)
	inQueue := data.InQueue(id)
	err := json.NewEncoder(responseWriter).Encode(inQueue)
	if err != nil {
		queueHandler.logger.Println("[ERROR] serializing queue", err)
		http.Error(responseWriter, "Unable to marshal json", http.StatusInternalServerError)
	}
}
