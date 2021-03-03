package data

import "testing"

func TestChecksValidation(t *testing.T) {
	player := &Player{
		UserID: 23,
	}

	err := player.ValidatePlayer()

	if err != nil {
		t.Fatal(err)
	}
}

func TestPlayerAlreadyInQueue(t *testing.T) {
	player := &Player{
		UserID: 42,
	}

	err := player.ValidatePlayer()

	if !(err != nil && err.Error() == "Key: 'Player.UserID' Error:Field validation for 'UserID' failed on the 'notinqueue' tag") {
		t.Fatal("The player is assumed to be already in the queue and should not be added again")
	}
}
