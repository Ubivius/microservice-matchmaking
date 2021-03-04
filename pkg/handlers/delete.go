package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-matchmaking/pkg/data"
)

// Delete a player with specified id from the queue
func (queueHandler *QueueHandler) Delete(responseWriter http.ResponseWriter, request *http.Request) {
	id := getPlayerID(request)
	queueHandler.logger.Println("Handle DELETE player", id)

	err := data.DeletePlayer(id)
	if err == data.ErrorPlayerNotFound {
		queueHandler.logger.Println("[ERROR] deleting, id does not exist")
		http.Error(responseWriter, "Player not found", http.StatusNotFound)
		return
	}

	if err != nil {
		queueHandler.logger.Println("[ERROR] deleting player", err)
		http.Error(responseWriter, "Error deleting player", http.StatusInternalServerError)
		return
	}

	responseWriter.WriteHeader(http.StatusNoContent)
}
