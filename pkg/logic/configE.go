package daw

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	db "github.com/grayson40/daw/pkg/db"
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
	user := db.CreateUser(username, email)

	// Add user to db
	err := db.AddUser(user)
	if err != nil {
		log.Fatal(err)
	}

	// Add credentials to json file
	file, err := json.MarshalIndent(user, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("./.daw/credentials.json", file, 0644)
}
