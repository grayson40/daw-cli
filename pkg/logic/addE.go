/*
Copyright © 2022 Grayson Crozier <grayson40@gmail.com>
*/
package daw

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	var inFiles []types.File

	// TODO: get tracked files
	// trackedFiles := GetTrackedFiles()

	if input[0] == "." {
		path, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		dirFiles, err := ioutil.ReadDir(path)
		if err != nil {
			panic(err)
		}

		for _, file := range dirFiles {
			name := file.Name()
			splitString := strings.Split(name, ".")
			if splitString[1] != "flp" {
				continue
			}
			path, err := filepath.Abs(name)
			if err != nil {
				panic(err)
			}
			modTime := GetModifiedTime(name)
			inFiles = append(inFiles, types.File{
				Name:  name,
				Path:  path,
				Saved: modTime,
			})
		}
	} else {
		for _, file := range input {
			name := file
			path, err := filepath.Abs(file)
			if err != nil {
				panic(err)
			}
			inFiles = append(inFiles, types.File{Name: name, Path: path})
		}
	}

	err := writeStaged(inFiles)
	if err != nil {
		panic(err)
	}
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
