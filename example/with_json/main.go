package main

import (
	"encoding/json"
	"fmt"

	"github.com/alaa-aqeel/optional-value"
)

type Player struct {
	Name     optional.Optional[string]   `json:"name"`
	Age      optional.Optional[int64]    `json:"age"`
	Chats    optional.Optional[[]string] `json:"chats"`
	IsOnline optional.Optional[bool]     `json:"is_online"`
	IsInGame optional.Optional[bool]     `json:"is_in_game"`
}

func main() {
	jsonInput := []byte(`{
		"name": "test", 
		"age": 20, 
		"is_online": false, 
		"chats": [
			"hi", 
			"how are you?"
		]
	}`)
	var player Player
	json.Unmarshal(jsonInput, &player)

	if player.IsInGame.IsSet {
		fmt.Printf("Player %s in game: %t \n ", player.Name.Value, player.IsInGame.Value)
	}

	if player.IsOnline.IsSet {
		fmt.Printf("Player %s is online: %t \n ", player.Name.Value, player.IsOnline.Value)
		// Player test is online: false
	}

	if player.Chats.IsSet && !player.Chats.IsEmpty() {
		fmt.Printf("Player %s chat: %v \n ", player.Name.Value, player.Chats.Value)
		// output: Player test chat: [hi how are you?]
	}

	fmt.Println(player)
	//  {
	// 		{test true }
	// 		{20 true 0}
	// 		{[hi how are you?] true []}
	// 		{false true false}
	// 		{false false false}
	// 	}
}
