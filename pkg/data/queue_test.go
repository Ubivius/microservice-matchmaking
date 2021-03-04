package data

import "testing"

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

func TestStartGame(t *testing.T) {
	player1 := &Player{
		UserID: 1,
		UserIP: "123.123.123.123",
	}
	player2 := &Player{
		UserID: 2,
		UserIP: "123.123.123.123",
	}

	err := AddPlayer(player1)
	err = AddPlayer(player2)

	if err != nil {
		t.Fatal(err)
	}
}
