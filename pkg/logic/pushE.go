/*
Copyright © 2022 Grayson Crozier <grayson40@gmail.com>
*/
package daw

import (
	"fmt"
	"os"

	"github.com/grayson40/daw/pkg/db"
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

	// Push commits up local branch
	pushToBranch(commits)

	// Clear commits
	if err := os.Truncate("./.daw/commits.json", 0); err != nil {
		panic(err)
	}
}

func pushToBranch(commits []types.Commit) {
	// Update commits
	db.UpdateCommits(commits)
}
