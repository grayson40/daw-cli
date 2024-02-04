/*
Copyright © 2022 Grayson Crozier <grayson40@gmail.com>
*/
package daw

import (
	"fmt"
	"os"

	constants "github.com/grayson40/daw/constants"
	io "github.com/grayson40/daw/pkg/io"
)

func ExecuteInit() {
	// Create .daw directory and subdirectories
	io.MakeDir(constants.DawDir)
	io.MakeDir(constants.InfoDir)
	io.MakeDir(constants.ObjectsDir)

	// Create json files to store staged, tracked, commits
	io.CreateEmptyFile(constants.StagedPath)
	io.CreateEmptyFile(constants.TrackedPath)
	io.CreateEmptyFile(constants.CommitsPath)

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("Initialized empty Daw repository in " + path)
	fmt.Println("  (use \"daw config --username <username> --email <email>\" to configure user credentials)")
}
