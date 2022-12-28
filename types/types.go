/*
Copyright © 2022 Grayson Crozier <grayson40@gmail.com>
*/
package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Change struct {
	Message string `bson:"message,omitempty" json:"message"`
}

type File struct {
	Name  string    `bson:"name,omitempty" json:"name"`
	Path  string    `bson:"path,omitempty" json:"path"`
	Saved time.Time `bson:"saved,omitempty" json:"saved"`
}

type Project struct {
	Name    string    `bson:"name,omitempty" json:"name"`
	Path    string    `bson:"path,omitempty" json:"path"`
	Saved   time.Time `bson:"saved,omitempty" json:"saved"`
	Changes []Change  `bson:"changes,omitempty" json:"changes"`
}

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Email    string             `bson:"email,omitempty" json:"email"`
	UserName string             `bson:"username,omitempty" json:"username"`
	Projects []Project          `bson:"projects,omitempty" json:"projects"`
}

type UserID struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
}
