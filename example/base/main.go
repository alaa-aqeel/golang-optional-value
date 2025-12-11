package main

import (
	"fmt"

	"github.com/alaa-aqeel/optional-value"
)

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
