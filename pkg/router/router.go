package router

import (
	"log"
	"net/http"

	"github.com/Ubivius/microservice-matchmaking/pkg/handlers"
	"github.com/gorilla/mux"
)

// New Mux route handling with gorilla/mux
func New(queueHandler *handlers.QueueHandler, logger *log.Logger) *mux.Router {
	// Mux route handling with gorilla/mux
	router := mux.NewRouter()

	// Get Router
	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/queue/{id:[0-9]+}", queueHandler.InQueue)

	// Post router
	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/queue", queueHandler.AddPlayer)
	// postRouter.HandleFunc("/queue/lobby", queueHandler.AddPlayers)
	postRouter.Use(queueHandler.MiddlewarePlayerValidation)

	// Delete router
	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/queue/{id:[0-9]+}", queueHandler.Delete)
	// deleteRouter.HandleFunc("/queue/lobby", queueHandler.DeletePlayers)

	return router
}
