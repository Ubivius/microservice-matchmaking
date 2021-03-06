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
	case data.ErrorUserNotFound:
		log.Error(err, "UserID doesn't exist")
		http.Error(responseWriter, "UserID doesn't exist", http.StatusBadRequest)
		return
	case data.ErrorAlreadyInQueue:
		log.Error(err, "Player is already in queue")
		http.Error(responseWriter, "Player is already in queue", http.StatusBadRequest)
		return
	default:
		log.Error(err, "Error adding player")
		http.Error(responseWriter, "Error adding player", http.StatusInternalServerError)
		return
	}
}
