/*
Copyright © 2022 Grayson Crozier <grayson40@gmail.com>
*/
package daw

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/grayson40/daw/types"
)

// Write input files to staged file
func ExecuteAdd(input []string) {
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

	// Get staged files
	stagedFiles := GetStaged()

	// Get tracked files
	trackedFiles, err := GetTracked()

	// Add all local dir files to staging area
	if input[0] == "." {
		// Get path
		path, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		// Read all files in local dir
		dirFiles, err := ioutil.ReadDir(path)
		if err != nil {
			panic(err)
		}

		// Iterate over local dir files
		for _, file := range dirFiles {
			// Get file name
			name := file.Name()

			// Only want project files
			splitString := strings.Split(name, ".")
			if splitString[1] != "flp" {
				continue
			}

			// Get absolute file path
			path, err := filepath.Abs(name)
			if err != nil {
				log.Fatalf("fatal: pathspec '%s' did not match any files", name)
				return
			}

			// Get last modified time
			modTime := GetModifiedTime(name)

			// Add to tracked if untracked
			if !IsTrackedFile(name) {
				trackedFiles = append(trackedFiles, types.File{
					Name:  name,
					Path:  path,
					Saved: modTime,
				})
			}

			// Append file for staging
			if !isStaged(name) {
				stagedFiles = append(stagedFiles, types.File{
					Name:  name,
					Path:  path,
					Saved: modTime,
				})
			}
		}
	} else {
		// Iterate over inputted files
		for _, file := range input {
			// Get file name
			name := file

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

				// Add to tracked if untracked
				if !IsTrackedFile(file) {
					trackedFiles = append(trackedFiles, types.File{
						Name:  name,
						Path:  path,
						Saved: modTime,
					})
				}

				// Append file for staging
				if !isStaged(name) {
					stagedFiles = append(stagedFiles, types.File{
						Name:  name,
						Path:  path,
						Saved: modTime,
					})
				}
			}
		}
	}

	// Write to tracked json
	err = writeTracked(trackedFiles)
	if err != nil {
		panic(err)
	}

	// Write to staged json
	err = writeStaged(stagedFiles)
	if err != nil {
		panic(err)
	}
}

// Returns true if file is already staged
func isStaged(fileName string) bool {
	stagedFiles := GetStaged()

	for _, file := range stagedFiles {
		if file.Name == fileName {
			return true
		}
	}

	return false
}

// Writes commit array to json file, returns err
func writeStaged(files []types.File) error {
	file, err2 := json.MarshalIndent(files, "", "\t")
	if err2 != nil {
		panic(err2)
	}

	err := ioutil.WriteFile("./.daw/staged.json", file, 0644)

	return err
}

// Write tracked array to json file
func writeTracked(files []types.File) error {
	file, err := json.MarshalIndent(files, "", "\t")
	if err != nil {
		panic(err)
	}

	writeErr := ioutil.WriteFile("./.daw/tracked.json", file, 0644)

	return writeErr
}

// func runPythonScript() {
// 	var cmd *exec.Cmd

// 	if len(input) == 1 {
// 		// Execute python script
// 		cmd = exec.Command("python", "C:/Users/grays/src/repos/daw/pkg/scripts/parse-fl.py", "--input", input[0])
// 	} else {
// 		// Format file name args
// 		var files = strings.Join(input, " ")

// 		// Execute python script
// 		cmd = exec.Command("python", "C:/Users/grays/src/repos/daw/pkg/scripts/parse-fl.py", "--input", files)
// 	}

// 	stdout, err := cmd.StdoutPipe()
// 	if err != nil {
// 		panic(err)
// 	}

// 	stderr, err := cmd.StderrPipe()
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = cmd.Start()
// 	if err != nil {
// 		panic(err)
// 	}

// 	go copyOutput(stdout)
// 	go copyOutput(stderr)

// 	cmd.Wait()
// }

// func copyOutput(r io.Reader) {
// 	scanner := bufio.NewScanner(r)
// 	for scanner.Scan() {
// 		fmt.Println(scanner.Text())
// 	}
// }
