/*
Copyright © 2022 Grayson Crozier <grayson40@gmail.com>
*/
package daw

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/grayson40/daw/pkg/requests"
	"github.com/grayson40/daw/types"
)

func ExecutePush() {
	// Throw error if not an initialized repo
	if !IsInitialized() {
		fmt.Println("fatal: not a daw repository (or any of the parent directories): .daw")
		return
	}

	// Throw error if user credentials not configured
	if _, err := os.Stat("./.daw/credentials.json"); err != nil {
		fmt.Println("fatal: user credentials not configured\n  (use \"daw config --username <username> --email <email>\" to configure user credentials)")
		return
	}

	// Read committedProject
	// TODO: better way to do this
	committedProject := GetCommittedProject()
	if committedProject.Name == "" {
		fmt.Println("Everything up-to-date")
		return
	}

	// Get current user id
	userId := GetCurrentUser().ID.Hex()

	// if project exists, prepend changes; else, add project to db
	if projectExistsInDb(committedProject.Path, userId) {
		// Get project changes from db
		project, _ := requests.GetProjectByPath(committedProject.Path, userId)
		projectChanges := project.Changes

		// Prepend incoming changes
		updatedChanges := append(committedProject.Changes, projectChanges...)

		// Update changes
		updateProjectChanges(project.Name, updatedChanges, userId)
	} else {
		// Add project to db
		addProjectToDb(committedProject, userId)
	}

	// userProjects := requests.GetProjects(userId)
	// for _, userProject := range userProjects {
	// 	if userProject.Path == commits.Path {
	// 		// Get list of changes from db
	// 		// Prepend to list of changes
	// 		userProject.Changes = append(commits.Changes, userProject.Changes...)
	// 	}
	// }

	// Clear commits
	if err := os.Truncate("./.daw/commits.json", 0); err != nil {
		panic(err)
	}
}

// Makes request to update project changes
func updateProjectChanges(projectName string, projectChanges []types.Change, userId string) {
	// Send put request with updated changes list
	requests.UpdateChanges(projectName, projectChanges, userId)
}

// Adds project to db
func addProjectToDb(commits types.Project, userId string) {
	// Update commits
	requests.AddProject(commits, userId)
}

// Returns the current user's credentials
func GetCurrentUser() types.User {
	var user types.User

	jsonFile, err := os.Open("./.daw/credentials.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &user)

	return user
}

// Returns true if project exists in db
func projectExistsInDb(projectPath string, userId string) bool {
	_, exists := requests.GetProjectByPath(projectPath, userId)
	return exists
}
