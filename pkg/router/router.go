package router

import (
	"net/http"

	"github.com/Ubivius/microservice-matchmaking/pkg/handlers"
	tokenValidation "github.com/Ubivius/shared-authentication/pkg/auth"
	"github.com/gorilla/mux"
)

// Mux route handling with gorilla/mux
func New(queueHandler *handlers.QueueHandler) *mux.Router {
	log.Info("Starting router")
	router := mux.NewRouter()

	// Get Router
	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.Use(tokenValidation.Middleware)
	getRouter.HandleFunc("/queue/{id:[0-9a-z-]+}", queueHandler.InQueue)

	//Health Check
	healthRouter := router.Methods(http.MethodGet).Subrouter()
	healthRouter.HandleFunc("/health/live", queueHandler.LivenessCheck)
	healthRouter.HandleFunc("/health/ready", queueHandler.ReadinessCheck)

	// Post router
	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.Use(tokenValidation.Middleware)
	postRouter.HandleFunc("/queue", queueHandler.AddPlayer)
	// postRouter.HandleFunc("/queue/lobby", queueHandler.AddPlayers)
	postRouter.Use(queueHandler.MiddlewarePlayerValidation)

	// Delete router
	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.Use(tokenValidation.Middleware)
	deleteRouter.HandleFunc("/queue/{id:[0-9a-z-]+}", queueHandler.Delete)
	// deleteRouter.HandleFunc("/queue/lobby", queueHandler.DeletePlayers)

	return router
}
