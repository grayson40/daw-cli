package daw

import (
	constants "github.com/grayson40/daw/constants"
	io "github.com/grayson40/daw/pkg/io"
)

// Returns true if user credentials are configured, false otherwise
func UserConfigured() bool {
	// Return early if credentials file does not exist
	if !io.FileExists(constants.CredentialsPath) {
		return false
	}

	userBytes := io.ReadFile(constants.CredentialsPath)

	// If username and email are empty, return false
	if len(userBytes) == 0 {
		return false
	}

	return true
}
