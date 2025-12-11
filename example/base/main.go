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
	Email    optional.Optional[string] `json:"email"`
}

func main() {

	user := User{
		Name:     optional.SetValue("Alaa Aqeel"),
		Age:      optional.SetValue[int64](29),
		Username: optional.SetValue(" "),
		IsActive: optional.SetValue(false),
		Email:    optional.NilValue[string](),
	}

	fmt.Println(user.Name.IsSet)                    // true
	fmt.Println(user.Name.Value)                    // "Alaa Aqeel"
	fmt.Println(user.IsActive.ValueOrDefault(true)) // false
	fmt.Println(user.Age.IsZero())                  // false
	fmt.Println(user.Username.IsEmpty())            // true
	fmt.Println(user.Email.IsSet)                   // false
}
