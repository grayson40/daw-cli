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
	// Read commits
	commits := ReadCommits()

	// Push commits up local branch
	pushToBranch(commits)

	// Clear commits
	if err := os.Truncate("commits.json", 0); err != nil {
		panic(err)
	}
}

func pushToBranch(commits []types.Commit) {
	// Push up branch linked to commits
	fmt.Print("pushing: ")
	fmt.Println(commits)
}
