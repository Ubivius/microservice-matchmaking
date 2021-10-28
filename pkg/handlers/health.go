package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-matchmaking/pkg/data"
)

// LivenessCheck determine when the application needs to be restarted
func (queueHandler *QueueHandler) LivenessCheck(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.WriteHeader(http.StatusOK)
}

//ReadinessCheck verifies that the application is ready to accept requests
func (queueHandler *QueueHandler) ReadinessCheck(responseWriter http.ResponseWriter, request *http.Request) {
	readinessProbeMicroserviceUser := data.MicroserviceUserPath + "/health/ready"

	_, errMicroserviceUser := http.Get(readinessProbeMicroserviceUser)

	if errMicroserviceUser != nil {
		log.Error(errMicroserviceUser, "Microservice-user unavailable")
		http.Error(responseWriter, "Microservice-user unavailable", http.StatusServiceUnavailable)
		return
	}

	responseWriter.WriteHeader(http.StatusOK)
}
