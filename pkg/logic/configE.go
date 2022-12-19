/*
Copyright © 2022 Grayson Crozier <grayson40@gmail.com>
*/
package daw

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	req "github.com/grayson40/daw/pkg/requests"
	"github.com/grayson40/daw/types"
)

func ExecuteConfig(username string, email string) {
	// Throw error if not an initialized repo
	if !IsInitialized() {
		fmt.Println("fatal: not a daw repository (or any of the parent directories): .daw")
		return
	}

	// Throw error if user credentials already configured
	if _, err := os.Stat("./.daw/credentials.json"); err == nil {
		fmt.Println("fatal: user credentials already configured")
		return
	}

	// Create user
	user := createUser(email, username)

	// Post user to db if not in there
	if !userExists(user) {
		user.ID = req.AddUser(user)
	}

	// Add credentials to json file
	writeUserCredentials(user)
}

// Writes user credentials to json
func writeUserCredentials(user types.User) {
	// Add credentials to json file
	file, err := json.MarshalIndent(user, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("./.daw/credentials.json", file, 0644)
}

// Returns user type
func createUser(email string, username string) types.User {
	return types.User{
		Email:    email,
		UserName: username,
		Projects: []types.Project{},
	}
}

// Returns true if user exists in db
func userExists(inUser types.User) bool {
	// See if user already exists
	users := req.GetUsers()
	for _, user := range users {
		if user.Email == inUser.Email && user.UserName == inUser.UserName {
			return true
		}
	}
	return false
}
