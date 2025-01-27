package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var exampleComRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@example\.com$`)

func ReadEmailsFile() []byte {
	fileName := "emails.json"

	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return data
}

func main() {
	var users []User

	data := ReadEmailsFile()

	err := json.Unmarshal(data, &users)

	if err != nil {
		panic(err)
	}

	var filteredUsers []User

	for _, user := range users {
		if exampleComRegex.Match([]byte(user.Email)) {
			filteredUsers = append(filteredUsers, user)
		}
	}

	fmt.Println("Struct is:", filteredUsers)

}
