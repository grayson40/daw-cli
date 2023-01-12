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
	"time"

	"github.com/grayson40/daw/types"
)

// POST request to append project
func AddProject(project types.Project, userId string) {
	// Encode the data
	postBody, _ := json.Marshal(project)
	responseBody := bytes.NewBuffer(postBody)

	// Make post request with project data
	REQUEST_URL := BASE_URL + "/projects?user_id=" + userId
	resp, err := http.Post(REQUEST_URL, "application/json", responseBody)
	if err != nil {
		log.Fatal(err.Error())
	}
	resp.Body.Close()
}

// GET request to get user projects
func GetProjects(userId string) []types.Project {
	// Make get request to get projects
	REQUEST_URL := BASE_URL + "/projects?user_id=" + userId
	resp, err := http.Get(REQUEST_URL)
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

// PUT request to update project changes in db
func UpdateChanges(projectName string, changes []types.Change, modTime time.Time, userId string) {
	// Encode the data
	postBody, _ := json.Marshal(map[string]interface{}{
		"changes": changes,
		"saved":   modTime,
	})
	responseBody := bytes.NewBuffer(postBody)

	// Make put request with updated changes
	client := &http.Client{}
	REQUEST_URL := BASE_URL + "/projects?user_id=" + userId + "&project_name=" + projectName
	req, err := http.NewRequest(http.MethodPut, REQUEST_URL, responseBody)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()
}

// Returns project and true if project exists, empty project and false otherwise
func GetProjectByPath(projectPath string, userId string) (types.Project, bool) {
	userProjects := GetProjects(userId)
	for _, userProject := range userProjects {
		if userProject.Path == projectPath {
			return userProject, true
		}
	}
	return types.Project{}, false
}

// func UpdateProjectByName(project types.Project, userId string) {
// 	// Encode the data
// 	postBody, _ := json.Marshal(project)
// 	responseBody := bytes.NewBuffer(postBody)

// 	// TODO: change this to a put request, updating projects commits and saved time
// 	// Make post request with project data
// 	resp, err := http.Post(BASE_URL+"/projects?id="+userId, "application/json", responseBody)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	// Doesn't really need to defer
// 	defer resp.Body.Close()
// }
