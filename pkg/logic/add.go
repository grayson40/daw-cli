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
	api "github.com/grayson40/daw/pkg/requests"
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

	// Get local staged projects
	stagedProjects := GetStagedProjects()

	// If input is ".", get all project files in directory
	if input[0] == "." {
		// Get all project files in directory
		files, err := filepath.Glob("*.flp")
		if err != nil {
			log.Fatalf(err.Error())
		}
		input = files
	}

	// Loop over the list of input files
	for _, projectFile := range input {
		name := projectFile

		// Only want project files
		splitString := strings.Split(name, ".")
		if splitString[1] != "flp" {
			fmt.Printf("fatal: pathspec '%s' is not valid for tracking", name)
			continue
		}

		// Check if file exists
		if _, err := os.Stat(name); err != nil {
			fmt.Printf("fatal: pathspec '%s' did not match any files", name)
			continue
		}

		// Get absolute file path
		path, err := filepath.Abs(name)
		if err != nil {
			log.Fatalf(err.Error())
		}

		// Diff project with latest commit
		diff := api.DiffProject(name)

		// If no diff, continue
		if diff == 0 {
			continue
		}

		// Get last modified time
		modTime := GetModifiedTime(name)

		file := types.File{
			Name:  name,
			Path:  path,
			Saved: modTime,
		}
		change := types.Change{
			Category:    "Low Pass",
			Instrument:  "Piano",
			Description: "Decreased by 10",
		}
		stagedProject := types.Project{
			Name:    name,
			File:    file,
			Changes: []types.Change{change},
		}

		// Flag to check if project is updated
		updated := false

		// Loop over the stagedProjects
		for i, project := range stagedProjects {
			// Check if the file path exists in the list
			if project.File.Path == stagedProject.File.Path {
				stagedProjects[i] = stagedProject
				updated = true
				break
			}
		}

		// If the project was not updated, append it to the list
		if !updated {
			stagedProjects = append(stagedProjects, stagedProject)
		}
	}

	// Write to staged json
	err := writeStaged(stagedProjects)
	if err != nil {
		panic(err)
	}
}

// Writes commit array to json file, returns error
func writeStaged(stagedProjects []types.Project) error {
	stagedObject, err := json.MarshalIndent(stagedProjects, "", "\t")
	if err != nil {
		panic(err)
	}

	io.WriteToFile(constants.StagedPath, stagedObject)

	return err
}
