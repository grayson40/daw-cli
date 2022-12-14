/*
Copyright © 2022 Grayson Crozier <grayson40@gmail.com>
*/
package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/grayson40/daw/types"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GLOBALS
var client *mongo.Client
var usersCollection *mongo.Collection

func init() {
	// Load env variables
	err := godotenv.Load("../local.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	// Configuration
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(os.Getenv("MONGODB_URL")).
		SetServerAPIOptions(serverAPIOptions)

	// Get db context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to db instance
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	// Grab users collection ref
	usersCollection = client.Database("daw").Collection("users")
}

// Returns user type
func CreateUser(username string, email string) types.User {
	user := types.User{
		UserName: username,
		Email:    email,
	}
	return user
}

// Adds user to mongoDB users collection. Returns error
func AddUser(user types.User) error {
	if UserExists(user.UserName) {
		return nil
	}
	_, err := usersCollection.InsertOne(context.TODO(), &user)
	return err
}

// Returns true if user exists, false otherwise
func UserExists(username string) bool {
	users := GetUsers()

	// Iterate through users
	for _, user := range users {
		if user["user_name"] == username {
			return true
		}
	}

	return false
}

// Returns a list of users in bson format
func GetUsers() []bson.M {
	// retrieve all the documents in users collection
	cursor, err := usersCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	// convert the cursor result to bson
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	return results
}

// func main() {

// 	// FOR QUERYING

// 	// retrieve all the documents in a collection
// 	cursor, err := usersCollection.Find(context.TODO(), bson.D{})
// 	if err != nil {
// 		log.Fatalf("Some error occured. Err: %s", err)
// 	}

// 	// convert the cursor result to bson
// 	var results []bson.M
// 	if err = cursor.All(context.TODO(), &results); err != nil {
// 		log.Fatalf("Some error occured. Err: %s", err)
// 	}

// 	// display the documents retrieved
// 	for _, result := range results {
// 		fmt.Println(result)
// 	}
// }
