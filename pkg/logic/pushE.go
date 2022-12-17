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

	// Read commits
	commits := GetCommits()
	if len(commits) == 0 {
		fmt.Println("Everything up-to-date")
		return
	}

	// // Append tracked files to db
	// trackedFiles, err := GetTracked()
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// Get current user id
	userId := GetCurrentUser().ID.Hex()
	currentUser := requests.GetUser(userId)

	// Get current user projects
	currentUserProjects := currentUser.Projects
	for _, commit := range commits {
		for _, file := range commit.Files {
			project := types.Project{
				File:    file,
				Commits: nil,
			}
			currentUserProjects = append(currentUserProjects, project)
		}
	}

	// Push commits up local branch
	// pushToBranch(commits)

	// Clear commits
	if err := os.Truncate("./.daw/commits.json", 0); err != nil {
		panic(err)
	}
}

// func pushToBranch(commits []types.Commit) {
// 	// Update commits
// 	db.UpdateCommits(commits)
// }

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
