/*
Copyright © 2022 Grayson Crozier <grayson40@gmail.com>
*/
package daw

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	constants "github.com/grayson40/daw/constants"
	io "github.com/grayson40/daw/pkg/io"
	"github.com/grayson40/daw/types"
)

// Write input files to staged file
func ExecuteAdd(input []string) {
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

	// Throw error if more than one file is inputted for staging
	if len(input) > 1 {
		fmt.Println("fatal: only one project file can be added at a time")
		return
	}

	// // Get current user id
	// userId := GetCurrentUser().ID.Hex()

	// Turn this into api call
	// Get staged project
	stagedProject := GetStagedProject()

	// Get project file input
	projectFile := input[0]

	// Get file name
	name := projectFile

	// Only want project files
	splitString := strings.Split(name, ".")
	if splitString[1] != "flp" {
		fmt.Printf("fatal: pathspec '%s' is not valid for tracking", name)
		return
	}

	// Check if file exists
	if _, err := os.Stat(name); err != nil {
		fmt.Printf("fatal: pathspec '%s' did not match any files", name)
	} else {
		// Get absolute file path
		path, err := filepath.Abs(name)
		if err != nil {
			log.Fatalf(err.Error())
		}

		// Get last modified time
		modTime := GetModifiedTime(name)

		// Append file for staging
		if !isStaged(path) {
			var changes []types.Change
			stagedProject = types.Project{
				Name:    name,
				Path:    path,
				Saved:   modTime,
				Changes: changes,
			}
		}
	}

	// Turn this into api call
	// Write to staged json
	err := writeStaged(stagedProject)
	if err != nil {
		panic(err)
	}
}

// Returns true if file is already staged
func isStaged(filepath string) bool {
	stagedProject := GetStagedProject()
	if stagedProject.Path == filepath {
		return true
	}
	return false
}

// Make this an api call
// Writes commit array to json file, returns err
func writeStaged(stagedProject types.Project) error {
	file, err := json.MarshalIndent(stagedProject, "", "\t")
	if err != nil {
		panic(err)
	}

	io.WriteToFile(constants.StagedPath, file)

	return err
}
