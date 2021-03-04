package data

import (
	"fmt"
)

// ErrorPlayerNotFound : Player specific errors
var ErrorPlayerNotFound = fmt.Errorf("Player not found")

// Player defines the structure for an player in queue.
// Formatting done with json tags to the right. "-" : don't include when encoding to json
type Player struct {
	UserID          int     `json:"userid" validate:"required,notinqueue,exist"`
	UserIP 	        string 	`json:"userip" validate:"required,ip"`
}

// Players is a collection of Player
type Players []*Player

// InQueue returns a boolean that verifies that a player is queued
func InQueue(id int) bool {
	return findIndexByPlayerID(id) > 0
}

// AddPlayer append a player to the queue
func AddPlayer(player *Player) {
	queue = append(queue, player)
}

// DeletePlayer deletes the player with the given id from the queue
func DeletePlayer(id int) error {
	index := findIndexByPlayerID(id)
	if index == -1 {
		return ErrorPlayerNotFound
	}

	queue = append(queue[:index], queue[index+1:]...)

	return nil
}

// Returns the index of a player in the queue
// Returns -1 when no player is found
func findIndexByPlayerID(id int) int {
	for index, player := range queue {
		if player.UserID == id {
			return index
		}
	}
	return -1
}

var queue = []*Player{
	{
		UserID: 2452,
		UserIP: "127.0.0.1",
	},
	{
		UserID: 42,
		UserIP: "192.223.10.1",
	},
}
