package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-matchmaking/pkg/data"
)

// AddPlayer add the player to the queue from the received JSON
func (queueHandler *QueueHandler) AddPlayer(responseWriter http.ResponseWriter, request *http.Request) {
	log.Info("AddPlayer request")
	player := request.Context().Value(KeyPlayer{}).(*data.Player)

	err := data.AddPlayer(player)
	switch err {
	case nil:
		responseWriter.WriteHeader(http.StatusNoContent)
		return
	default:
		log.Error(err, "Error adding player")
		http.Error(responseWriter, "Error adding player", http.StatusInternalServerError)
		return
	}
}
