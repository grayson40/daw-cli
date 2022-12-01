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
func newCommit(files []types.File, message string) types.Commit {
	c := types.Commit{
		Files:   files,
		Message: message,
	}
	return c
}

// Executes the commiting process
func ExecuteCommit(message string) {
	// Read staged files
	files := ReadStaged()

	// Read commit stack
	commits := ReadCommits()

	// Create new commit and write to commit stack
	commit := newCommit(files, message)
	commits = append([]types.Commit{commit}, commits...)
	err := writeCommit(commits)
	if err != nil {
		panic(err)
	}

	// Clear staged files
	if err := os.Truncate("staged.json", 0); err != nil {
		panic(err)
	}
}

// Reads contents of json staged files and returns array of staged files
func ReadStaged() []types.File {
	var files []types.File

	if _, err := os.Stat("staged.json"); err == nil {
		jsonFile, err := os.Open("staged.json")
		if err != nil {
			fmt.Println(err)
		}

		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &files)

		return files
	}

	return nil
}

// Reads contents of json commit file and returns array of commits
func ReadCommits() []types.Commit {
	var commits []types.Commit

	if _, err := os.Stat("commits.json"); err == nil {
		jsonFile, err := os.Open("commits.json")
		if err != nil {
			fmt.Println(err)
		}

		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &commits)

		return commits
	}

	return nil
}

// Writes commit array to json file, returns err
func writeCommit(commits []types.Commit) error {
	file, err2 := json.MarshalIndent(commits, "", "\t")
	if err2 != nil {
		panic(err2)
	}

	err := ioutil.WriteFile("commits.json", file, 0644)

	return err
}
