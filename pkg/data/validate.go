package data

import (
	"net"

	"github.com/go-playground/validator"
)

// ValidatePlayer a player with json validation and in queue validator
func (player *Player) ValidatePlayer() error {
	validate := validator.New()
	err1 := validate.RegisterValidation("exist", validateExist)
	err2 := validate.RegisterValidation("notinqueue", validateNotInQueue)
	err3 := validate.RegisterValidation("ip", validateIP)
	if err1 != nil {
		panic(err1)
	} else if err2 != nil {
		panic(err2)
	}else if err3 != nil {
		panic(err3)
	}

	return validate.Struct(player)
}

// validates that the player exist
func validateExist(fieldLevel validator.FieldLevel) bool {
	// validation of the UserID with a call to microservice-user 
	return true
}

// validates that the player is not already in the queue
func validateNotInQueue(fieldLevel validator.FieldLevel) bool {
	return !InQueue(int(fieldLevel.Field().Int()))
}

// Custom IP validator
func validateIP(fieldLevel validator.FieldLevel) bool {
	return net.ParseIP(fieldLevel.Field().String()) != nil
}
