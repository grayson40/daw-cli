/*
Copyright © 2022 Grayson Crozier <grayson40@gmail.com>
*/
package daw

import (
	"fmt"
	"os"

	"github.com/grayson40/daw/types"
)

func ExecutePush() {
	// Throw error if not an initialized repo
	if !IsInitialized() {
		fmt.Println("fatal: not a daw repository (or any of the parent directories): .daw")
		return
	}

	// Read commits
	commits := GetCommits()
	if len(commits) == 0 {
		fmt.Println("Everything up-to-date")
		return
	}

	// Push commits up local branch
	pushToBranch(commits)

	// Clear commits
	if err := os.Truncate("./.daw/commits.json", 0); err != nil {
		panic(err)
	}
}

func pushToBranch(commits []types.Commit) {
	// Push up branch linked to commits
	fmt.Print("pushing: ")
	fmt.Println(commits)
}
