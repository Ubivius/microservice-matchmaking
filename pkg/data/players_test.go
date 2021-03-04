package data

import "testing"

func TestChecksValidation(t *testing.T) {
	player := &Player{
		UserID: 23,
		UserIP: "123.123.123.123",
	}

	err := player.ValidatePlayer()

	if err != nil {
		t.Fatal(err)
	}
}

func TestInvalidIP(t *testing.T) {
	player := &Player{
		UserID: 1,
		UserIP: "423.123.123.123",
	}

	err := player.ValidatePlayer()

	if !(err != nil && err.Error() == "Key: 'Player.UserIP' Error:Field validation for 'UserIP' failed on the 'ip' tag") {
		t.Fatal("Each part of the IP address is assumed to be between 255 and 0")
	}
}

func TestPlayerAlreadyInQueue(t *testing.T) {
	player := &Player{
		UserID: 42,
		UserIP: "123.123.123.123",
	}

	err := player.ValidatePlayer()

	if !(err != nil && err.Error() == "Key: 'Player.UserID' Error:Field validation for 'UserID' failed on the 'notinqueue' tag") {
		t.Fatal("The player is assumed to be already in the queue and should not be added again")
	}
}

func TestAddPlayer(t *testing.T) {
	player := &Player{
		UserID: 1,
		UserIP: "123.123.123.123",
	}

	AddPlayer(player)

	if findIndexByPlayerID(player.UserID) == -1 {
		t.Fatal("The player has not been added to the queue")
	}
}

func TestDeletePlayer(t *testing.T) {
	err := DeletePlayer(42)

	if err != nil {
		t.Fatal(err)
	}
}

func TestDeletePlayerNotInQueue(t *testing.T) {
	err := DeletePlayer(2)

	if !(err != nil && err.Error() == "Player not found") {
		t.Fatal("The player is not in the queue, so it is assumed not to be found")
	}
}
