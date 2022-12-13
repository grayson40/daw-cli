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
	"time"

	"github.com/grayson40/daw/types"
)

// Colors
const Red = "\033[31m"
const Green = "\033[32m"
const White = "\033[97m"

func ExecuteStatus() {
	// Throw error if not an initialized repo
	if !IsInitialized() {
		fmt.Println("fatal: not a daw repository (or any of the parent directories): .daw")
		return
	}

	// Get staged files
	staged := GetStaged()

	// TODO: Build function to parse and display tracked files not staged for commit (red)
	notStaged := GetNotStaged()

	// Get untracked files
	notTracked := getUntracked()

	if len(staged) != 0 {
		// Show changed files to be committed (green)
		fmt.Println("Changes to be committed:\n  (use \"daw restore --staged <file>...\" to unstage)")
		for _, file := range staged {
			fmt.Println(Green + "\t" + file.Name + White)
		}
		// New line for formatting
		fmt.Println()
	} else {
		// Show no changes added if staged is empty
		defer fmt.Print("no changes added to commit (use \"daw add\" and/or \"daw commit <message>\")")
	}

	// Show changed files not staged for commit
	if len(notStaged) != 0 {
		fmt.Print("Changes not staged for commit:\n  (use \"daw add <file>...\" to update what will be committed)\n  (use \"daw restore <file>...\" to discard changes in working directory)\n")
		for _, file := range notStaged {
			fmt.Println("\t" + Red + file.Name + White)
		}
		fmt.Println()
	}

	// Display untracked files
	if len(notTracked) != 0 {
		fmt.Println("Untracked files:\n  (use \"daw add <file>...\" to include in what will be committed)")
		for _, file := range notTracked {
			fmt.Println("\t" + Red + file.Name + White)
		}
		fmt.Println()
	}
}

// Returns an array of untracked projects in working directory
func getUntracked() []types.File {
	var notTracked []types.File

	// Get working directory path
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Get files in working directory
	dirFiles, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	// Append working directory files
	for _, file := range dirFiles {
		fileExtension := filepath.Ext(file.Name())

		// Check if file is tracked
		if fileExtension == ".flp" && !IsTrackedFile(file.Name()) {
			notTracked = append(notTracked, types.File{Name: file.Name(), Path: path})
		}
	}

	return notTracked
}

// Returns an array of tracked files that are not staged
func GetNotStaged() []types.File {
	var notStaged []types.File

	// Get working directory path
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Get files in working directory
	dirFiles, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, file := range dirFiles {
		// If file is tracked, not staged and has changes. Append to list
		if IsTrackedFile(file.Name()) && !IsStagedFile(file.Name()) && isModifiedFile(file.Name()) {
			notStaged = append(notStaged, types.File{Name: file.Name(), Path: path})
		}
	}

	return notStaged
}

func isModifiedFile(fileName string) bool {
	// get last modified time from dir
	modTime := GetModifiedTime(fileName)

	// get last modified time from json
	trackedFiles, err := GetTracked()
	if err != nil {
		panic(err)
	}

	// Find tracked file and compare mod times
	for _, file := range trackedFiles {
		if file.Name == fileName {
			return modTime != file.Saved
		}
	}

	return false
}

// Returns last modified time from local dir
func GetModifiedTime(fileName string) time.Time {
	f, err := os.Stat(fileName)
	if err != nil {
		fmt.Println(err)
	}
	return f.ModTime()
}

// Returns true if file is staged, false otherwise
func IsStagedFile(fileName string) bool {
	stagedFiles := GetStaged()

	for _, file := range stagedFiles {
		if file.Name == fileName {
			return true
		}
	}

	return false
}

// Returns true if file is tracked, false otherwise
func IsTrackedFile(fileName string) bool {
	trackedFiles, err := GetTracked()
	if err != nil {
		return false
	}

	for _, file := range trackedFiles {
		if file.Name == fileName {
			return true
		}
	}

	return false
}

// Returns an array of tracked files
func GetTracked() ([]types.File, error) {
	var trackedFiles []types.File

	jsonFile, err := os.Open("./.daw/tracked.json")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &trackedFiles)

	return trackedFiles, err
}
