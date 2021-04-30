package data

import (
	"github.com/go-playground/validator"
)

// ValidatePlayer a player with json validation and in queue validator
func (player *Player) ValidatePlayer() error {
	validate := validator.New()

	return validate.Struct(player)
}
