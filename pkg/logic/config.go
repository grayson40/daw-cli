/*
Copyright © 2022 Grayson Crozier <grayson40@gmail.com>
*/
package daw

import (
	"encoding/json"
	"fmt"
	"log"

	constants "github.com/grayson40/daw/constants"
	io "github.com/grayson40/daw/pkg/io"
	api "github.com/grayson40/daw/pkg/requests"
	"github.com/grayson40/daw/types"
)

func ExecuteConfig(username string, email string) {
	// Throw error if not an initialized repo
	if !io.IsInitialized() {
		fmt.Println("fatal: not a daw repository (or any of the parent directories): .daw")
		return
	}

	// Throw error if user credentials already configured
	if UserConfigured() {
		fmt.Println("fatal: user credentials already configured")
		return
	}

	// Construct user
	user := constructUser(email, username)

	// Post user to db if not in there, get ID
	if !userExists(user) {
		user.ID = api.AddUser(user)
	} else {
		user.ID = api.GetUserIdByEmail(email)
	}

	// Add credentials to json file
	writeUserCredentials(user)
}

// Writes user credentials to json file
func writeUserCredentials(user types.User) {
	file, err := json.MarshalIndent(user, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	io.WriteToFile(constants.CredentialsPath, file)
}

// Returns populated user type
func constructUser(email string, username string) types.User {
	return types.User{
		Email:    email,
		UserName: username,
		Projects: []types.Project{},
	}
}

// Returns true if user exists in db
func userExists(inUser types.User) bool {
	// See if user already exists
	users := api.GetUsers()
	for _, user := range users {
		if user.Email == inUser.Email && user.UserName == inUser.UserName {
			return true
		}
	}
	return false
}
