/*
Copyright © 2022 Grayson Crozier <grayson40@gmail.com>
*/
package types

import "time"

type Commit struct {
	Files   []File `json:"Files"`
	Message string `json:"Message"`
	// Branch  string `json:"Branch"`
}

type File struct {
	Name  string    `json:"Name"`
	Path  string    `json:"Path"`
	Saved time.Time `json:"Saved"`
}
