/*
Copyright © 2022 Grayson Crozier <grayson40@gmail.com>
*/
package types

import (
	"time"
)

type File struct {
	Name     string
	Path     string
	Saved    time.Time
	Contents string
}

type Project struct {
	Name    string
	File    File
	Changes []Change
}

type User struct {
	ID       int
	Username string
	Email    string
}

type Commit struct {
	Message string
	Created time.Time
}

type Change struct {
	Instrument  string
	Category    string
	Description string
}
