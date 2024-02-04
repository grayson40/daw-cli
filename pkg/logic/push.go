/*
Copyright © 2022 Grayson Crozier <grayson40@gmail.com>
*/
package daw

import (
	"encoding/json"
	"fmt"
	"time"

	constants "github.com/grayson40/daw/constants"
	io "github.com/grayson40/daw/pkg/io"
	api "github.com/grayson40/daw/pkg/requests"
	"github.com/grayson40/daw/types"
)

func ExecutePush() {
	// Throw error if not an initialized repo
	if !io.IsInitialized() {
		fmt.Println("fatal: not a daw repository (or any of the parent directories): .daw")
		return
	}

	// Throw error if user credentials not configured
	if !UserConfigured() {
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
		project, _ := api.GetProjectByPath(committedProject.Path, userId)
		projectChanges := project.Changes

		// Prepend incoming changes
		updatedChanges := append(committedProject.Changes, projectChanges...)

		// Get latest savedTime
		modTime := GetModifiedTime(project.Name)

		// Update changes
		updateProjectChanges(project.Name, updatedChanges, modTime, userId)
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
	clearCommits()
}

// Clears commits
func clearCommits() {
	io.ClearFile(constants.CommitsPath)
}

// Makes request to update project changes
func updateProjectChanges(projectName string, projectChanges []types.Change, modTime time.Time, userId string) {
	api.UpdateChanges(projectName, projectChanges, modTime, userId)
}

// Adds project to db
func addProjectToDb(commits types.Project, userId string) {
	api.AddProject(commits, userId)
}

// Returns the current user's credentials
func GetCurrentUser() types.User {
	var user types.User

	userBytes := io.ReadFile(constants.CredentialsPath)
	json.Unmarshal(userBytes, &user)

	return user
}

// Returns true if project exists in db
func projectExistsInDb(projectPath string, userId string) bool {
	_, exists := api.GetProjectByPath(projectPath, userId)
	return exists
}
