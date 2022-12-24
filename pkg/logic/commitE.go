/*
Copyright © 2022 Grayson Crozier <grayson40@gmail.com>
*/
package daw

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/grayson40/daw/types"
)

// Creates and returns a new commit
func newCommit(projects []types.Project, message string) []types.Project {
	c := types.Change{
		Message: message,
	}
	for _, project := range projects {
		project.Changes = append(project.Changes, c)
	}
	return projects
}

// Executes the commiting process
func ExecuteCommit(message string) {
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

	// Create change type
	change := types.Change{Message: message}

	// Get currently staged project
	stagedProject := GetStagedProject()

	// Throw error if no project staged
	// TODO: find better way to do this
	if stagedProject.Name == "" {
		fmt.Println("error: no changes added to commit")
		return
	}

	// Read commit stack
	committedProject := GetCommittedProject()

	// Get changes and append
	changes := committedProject.Changes
	changes = append([]types.Change{change}, changes...)
	stagedProject.Changes = changes

	// TODO: add commit(s) to designated project files
	// TODO: get project(s) & commits of project(s)
	// TODO: append commit(s) to designated project file
	// So commits.json will be a list of updated project files

	// Get current user projects
	// currentUserId := GetCurrentUser().ID.Hex()
	// currentUserProjects := requests.GetProjects(currentUserId)

	// if len(currentUserProjects) == 0 {
	// } else {
	// 	// Search for project and append changes to changes list
	// }

	// for _, stagedProject := range stagedProjects {
	// 	if !fileExists(currentUserProjects, stagedProject) {
	// 		project := requests.GetProjectByName(stagedProject.Name, currentUserId)
	// 		// Update project commits
	// 		for _, commit := range committedProjects {
	// 			project.Commits = append(stagedProject.Commits, commit.Commits...)
	// 		}
	// 		requests.UpdateProjectByName(project, currentUserId)
	// 	}
	// }

	// Create new commit and write to commit stack
	writeErr := writeCommit(stagedProject)
	if writeErr != nil {
		panic(writeErr)
	}

	// Clear staged files
	if err := os.Truncate("./.daw/staged.json", 0); err != nil {
		panic(err)
	}
}

// Reads contents of json staged files and returns array of staged files
func GetStagedProject() types.Project {
	var project types.Project

	jsonFile, err := os.Open("./.daw/staged.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	if byteValue == nil {
		return types.Project{}
	}

	json.Unmarshal(byteValue, &project)
	return project
}

// Reads contents of json commit file and returns array of commits
func GetCommittedProject() types.Project {
	var project types.Project

	jsonFile, err := os.Open("./.daw/commits.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &project)

	return project
}

// Writes commit array to json file, returns err
func writeCommit(committedProject types.Project) error {
	file, err := json.MarshalIndent(committedProject, "", "\t")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("./.daw/commits.json", file, 0644)

	return err
}

// Returns true if project file exists in db
func fileExists(currentUserProjects []types.Project, inProject types.Project) bool {
	for _, currentUserProject := range currentUserProjects {
		if inProject.Path == currentUserProject.Path {
			return true
		}
	}
	return false
}
