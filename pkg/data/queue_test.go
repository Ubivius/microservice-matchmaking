package data

import "testing"

func TestAddPlayer(t *testing.T) {
	player := &Player{
		UserID: 1,
		UserIP: "123.123.123.123",
	}

	err := AddPlayer(player)

	if err != nil {
		t.Fatal(err)
	}
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

func TestStartGame(t *testing.T) {
	player1 := &Player{
		UserID: 1,
		UserIP: "123.123.123.123",
	}
	player2 := &Player{
		UserID: 2,
		UserIP: "123.123.123.123",
	}

	errAddPlayer1 := AddPlayer(player1)
	errAddPlayer2 := AddPlayer(player2)

	if errAddPlayer1 != nil || errAddPlayer2 != nil {
		t.Fatal("Error adding players and sending the request to the Dispatcher")
	}
}
