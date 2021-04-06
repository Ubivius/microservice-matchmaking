package router

import (
	"net/http"

	"github.com/Ubivius/microservice-matchmaking/pkg/handlers"
	"github.com/gorilla/mux"
)

// Mux route handling with gorilla/mux
func New(queueHandler *handlers.QueueHandler) *mux.Router {
	log.Info("Starting router")
	router := mux.NewRouter()

	// Get Router
	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/queue/{id:[0-9a-z-]+}", queueHandler.InQueue)

	// Post router
	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/queue", queueHandler.AddPlayer)
	// postRouter.HandleFunc("/queue/lobby", queueHandler.AddPlayers)
	postRouter.Use(queueHandler.MiddlewarePlayerValidation)

	// Delete router
	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/queue/{id:[0-9a-z-]+}", queueHandler.Delete)
	// deleteRouter.HandleFunc("/queue/lobby", queueHandler.DeletePlayers)

	return router
}
