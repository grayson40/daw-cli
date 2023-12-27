/*
Copyright © 2022 Grayson Crozier <grayson40@gmail.com>
*/
package daw

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Colors
// const Red = "\033[31m"
// const Green = "\033[32m"
// const White = "\033[97m"

func ExecuteRestore(stagedFile string) {
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

	// Read the staged.json file
	data, err := ioutil.ReadFile("./.daw/staged.json")
	if err != nil {
		fmt.Println("Error reading staged.json:", err)
		return
	}

	// Unmarshal the JSON data into a map
	var stagedData map[string]interface{}
	err = json.Unmarshal(data, &stagedData)
	if err != nil {
		fmt.Println("Error unmarshalling staged.json:", err)
		return
	}

	// Delete the key corresponding to the filename
	delete(stagedData, stagedFile)

	// Marshal the map back into JSON
	data, err = json.Marshal(stagedData)
	if err != nil {
		fmt.Println("Error marshalling staged.json:", err)
		return
	}

	// Write the JSON data back to the staged.json file
	err = ioutil.WriteFile("./.daw/staged.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing staged.json:", err)
		return
	}
}
