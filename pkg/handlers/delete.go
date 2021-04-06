package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-matchmaking/pkg/data"
)

// Delete a player with specified id from the queue
func (queueHandler *QueueHandler) Delete(responseWriter http.ResponseWriter, request *http.Request) {
	id := getPlayerID(request)
	log.Info("Delete player by ID request", "id", id)

	err := data.DeletePlayer(id)
	switch err {
	case nil:
		responseWriter.WriteHeader(http.StatusNoContent)
		return
	case data.ErrorPlayerNotFound:
		log.Error(err, "Error deleting player, id does not exist")
		http.Error(responseWriter, "Player not found", http.StatusNotFound)
		return
	default:
		log.Error(err, "Error deleting player")
		http.Error(responseWriter, "Error deleting player", http.StatusInternalServerError)
		return
	}
}
