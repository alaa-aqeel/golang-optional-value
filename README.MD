# Optional Value 
This package provides a way to work with optional values in Go. It allows you to create, set and get values of any type in a safe way, without having to worry about nil pointer errors. It also provides a way to create values with default values, which can be useful when working with JSON data.


# Example
### Base Example 
```golang
type User struct {
	Name     optional.Optional[string] `json:"name"`
	Age      optional.Optional[int64]  `json:"age"`
	Username optional.Optional[string] `json:"usernmae"`
	IsActive optional.Optional[bool]   `json:"is_active"`
}

func main() {

	user := User{
		Name:     optional.SetValue("Alaa Aqeel"),
		Age:      optional.SetValue[int64](29),
		Username: optional.SetValue("username"),
		IsActive: optional.SetValue(false),
	}

	fmt.Println(user.Name.IsSet)                    // check if value is set
	fmt.Println(user.Name.Value)                    // get value
	fmt.Println(user.IsActive.ValueOrDefault(true)) // get value or default
	fmt.Println(user.Age.IsZero())                  // check if value is zero
	fmt.Println(user.Username.IsEmpty())            // check if value is empty
}

```

### with json 
```golang
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
```