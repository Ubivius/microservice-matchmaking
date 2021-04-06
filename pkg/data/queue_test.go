package data

import "testing"

func TestAddPlayer(t *testing.T) {
	player := &Player{
		UserID: "a2181017-5c53-422b-b6bc-036b27c04fc8",
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
	err := DeletePlayer("a2181017-5c53-422b-b6bc-036b27c04fc8")

	if err != nil {
		t.Fatal(err)
	}
}

func TestDeletePlayerNotInQueue(t *testing.T) {
	err := DeletePlayer("6df55d54-96f3-11eb-a8b3-0242ac130003")

	if !(err != nil && err.Error() == "Player not found") {
		t.Fatal("The player is not in the queue, so it is assumed not to be found")
	}
}

func TestStartGame(t *testing.T) {
	player1 := &Player{
		UserID: "6df55d54-96f3-11eb-a8b3-0242ac130003",
		UserIP: "123.123.123.123",
	}
	player2 := &Player{
		UserID: "73dfa062-96f3-11eb-a8b3-0242ac130003",
		UserIP: "123.123.123.123",
	}

	errAddPlayer1 := AddPlayer(player1)
	errAddPlayer2 := AddPlayer(player2)

	if errAddPlayer1 != nil || errAddPlayer2 != nil {
		t.Fatal("Error adding players and sending the request to the Dispatcher")
	}
}
