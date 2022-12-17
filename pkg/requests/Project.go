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

// POST request to append project
func AddProject(project types.Project, userId string) {
	// Encode the data
	postBody, _ := json.Marshal(project)
	responseBody := bytes.NewBuffer(postBody)

	// Make post request with project data
	resp, err := http.Post(BASE_URL+"/user?id="+userId+"/projects", "application/json", responseBody)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Doesn't really need to defer
	defer resp.Body.Close()
}

// GET request to get user projects
func GetProjects(userId string) []types.Project {
	// Response
	resp, err := http.Get(BASE_URL + "/user?id=" + userId + "/projects")
	if err != nil {
		log.Fatal(err.Error())
	}

	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Decode json response
	var projects []types.Project
	json.Unmarshal(body, &projects)

	return projects
}
