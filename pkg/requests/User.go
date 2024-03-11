/*
Copyright © 2022 Grayson Crozier <grayson40@gmail.com>
*/
package requests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/grayson40/daw/types"
)

// POST request to create user
func AddUser(user types.User) int {
	// Encode the data
	postBody, _ := json.Marshal(user)
	responseBody := bytes.NewBuffer(postBody)

	// Make post request with user data
	URL := BASE_URL + "/api/v1/users"
	resp, err := http.Post(URL, "application/json", responseBody)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Doesn't really need to defer
	defer resp.Body.Close()

	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Decode json response
	var createdID int
	json.Unmarshal(body, &createdID)

	return createdID
}

// GET request to get users
func GetUser(userId string) types.User {
	// Response
	resp, err := http.Get(BASE_URL + "/user?id=" + userId)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Decode json response
	var user types.User
	json.Unmarshal(body, &user)

	return user
}

// GET request to get users
func GetUsers() []types.User {
	// Response
	resp, err := http.Get(BASE_URL + "/api/v1/users")
	if err != nil {
		log.Fatal(err.Error())
	}

	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Decode json response
	var users []types.User
	json.Unmarshal(body, &users)

	return users
}

// GET request to get user id by email
func GetUserIdByEmail(email string) int {
	// Response
	resp, err := http.Get(BASE_URL + "/user?email=" + email)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Decode json response
	var user types.User
	json.Unmarshal(body, &user)

	return int(user.ID)
}
