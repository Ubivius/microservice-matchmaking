package data

import "testing"

func TestChecksValidation(t *testing.T) {
	player := &Player{
		UserID: "b70e1d8c-96f3-11eb-a8b3-0242ac130003",
		UserIP: "123.123.123.123",
	}

	err := player.ValidatePlayer()

	if err != nil {
		t.Fatal(err)
	}
}

func TestInvalidIP(t *testing.T) {
	player := &Player{
		UserID: "73dfa062-96f3-11eb-a8b3-0242ac130003",
		UserIP: "423.123.123.123",
	}

	err := player.ValidatePlayer()

	if !(err != nil && err.Error() == "Key: 'Player.UserIP' Error:Field validation for 'UserIP' failed on the 'ip' tag") {
		t.Fatal("Each part of the IP address is assumed to be between 255 and 0")
	}
}

func TestPlayerAlreadyInQueue(t *testing.T) {
	player := &Player{
		UserID: "a2181017-5c53-422b-b6bc-036b27c04fc8",
		UserIP: "123.123.123.123",
	}

	errAddPlayer := AddPlayer(player)

	if errAddPlayer != nil {
		t.Fatal(errAddPlayer)
	}

	err := player.ValidatePlayer()

	if !(err != nil && err.Error() == "Key: 'Player.UserID' Error:Field validation for 'UserID' failed on the 'notinqueue' tag") {
		t.Fatal("The player is assumed to be already in the queue and should not be added again")
	}
}
