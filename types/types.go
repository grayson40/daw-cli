/*
Copyright © 2022 Grayson Crozier <grayson40@gmail.com>
*/
package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Commit struct {
	Files   []File `bson:"files,omitempty" json:"files"`
	Message string `bson:"message,omitempty" json:"message"`
	// Branch  string `json:"Branch"`
}

type File struct {
	Name  string    `bson:"name,omitempty" json:"name"`
	Path  string    `bson:"path,omitempty" json:"path"`
	Saved time.Time `bson:"saved,omitempty" json:"saved"`
}

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Email    string             `bson:"email,omitempty" json:"email"`
	UserName string             `bson:"user_name,omitempty" json:"user_name"`
	Commits  []Commit           `bson:"commits,omitempty" json:"commits"`
}
