package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-matchmaking/pkg/data"
)

// AddPlayer add the player to the queue from the received JSON
func (queueHandler *QueueHandler) AddPlayer(responseWriter http.ResponseWriter, request *http.Request) {
	queueHandler.logger.Println("Handle POST Player")
	player := request.Context().Value(KeyPlayer{}).(*data.Player)

	data.AddPlayer(player)
	responseWriter.WriteHeader(http.StatusNoContent)
}
