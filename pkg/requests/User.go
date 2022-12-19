/*
Copyright © 2022 Grayson Crozier <grayson40@gmail.com>
*/
package requests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/grayson40/daw/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// POST request to create user
func AddUser(user types.User) primitive.ObjectID {
	// Encode the data
	postBody, _ := json.Marshal(user)
	responseBody := bytes.NewBuffer(postBody)

	// Make post request with user data
	resp, err := http.Post(BASE_URL+"/user", "application/json", responseBody)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Doesn't really need to defer
	defer resp.Body.Close()

	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Decode json response
	var userID types.UserID
	json.Unmarshal(body, &userID)

	return userID.ID
}

// GET request to get users
func GetUser(userId string) types.User {
	// Response
	resp, err := http.Get(BASE_URL + "/user?id=" + userId)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Decode json response
	var user types.User
	json.Unmarshal(body, &user)

	return user
}

// GET request to get users
func GetUsers() []types.User {
	// Response
	resp, err := http.Get(BASE_URL + "/users")
	if err != nil {
		log.Fatal(err.Error())
	}

	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Decode json response
	var users []types.User
	json.Unmarshal(body, &users)

	return users
}

// package db

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"os"
// 	"time"

// 	"github.com/grayson40/daw/types"
// 	"github.com/joho/godotenv"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// // GLOBALS
// var client *mongo.Client
// var usersCollection *mongo.Collection

// func init() {
// 	// Load env variables
// 	err := godotenv.Load("../local.env")
// 	if err != nil {
// 		log.Fatalf("Some error occured. Err: %s", err)
// 	}

// 	// Configuration
// 	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
// 	clientOptions := options.Client().
// 		ApplyURI(os.Getenv("MONGODB_URL")).
// 		SetServerAPIOptions(serverAPIOptions)

// 	// Get db context
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	// Connect to db instance
// 	client, err = mongo.Connect(ctx, clientOptions)
// 	if err != nil {
// 		log.Fatalf("Some error occured. Err: %s", err)
// 	}

// 	// Grab users collection ref
// 	usersCollection = client.Database("daw").Collection("users")
// }

// // Returns user type
// func CreateUser(username string, email string) types.User {
// 	user := types.User{
// 		UserName: username,
// 		Email:    email,
// 		Commits:  nil,
// 	}
// 	return user
// }

// // Adds user to mongoDB users collection. Returns error
// func AddUser(user types.User) error {
// 	if UserExists(user.Email) {
// 		return nil
// 	}
// 	_, err := usersCollection.InsertOne(context.TODO(), &user)
// 	return err
// }

// // Returns the current user
// func GetCurrentUser() types.User {
// 	var user types.User

// 	// Open json credentials file
// 	jsonFile, err := os.Open("./.daw/credentials.json")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	defer jsonFile.Close()

// 	// Read data into user
// 	byteValue, _ := ioutil.ReadAll(jsonFile)
// 	json.Unmarshal(byteValue, &user)

// 	// Find user by email and decode res from db
// 	var currentUser types.User
// 	err = usersCollection.FindOne(context.TODO(), bson.D{{
// 		Key:   "email",
// 		Value: user.Email,
// 	}}).Decode(&currentUser)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	return currentUser
// }

// // Returns current user
// func GetUserCommits() []types.Commit {
// 	return GetCurrentUser().Commits
// }

// // Pushes commits to db
// func UpdateCommits(commits []types.Commit) *mongo.UpdateResult {
// 	// Get current user
// 	currentUser := GetCurrentUser()

// 	// Append pushed commits
// 	for _, commit := range commits {
// 		currentUser.Commits = append(currentUser.Commits, commit)
// 	}

// 	// Grab user id
// 	id, _ := primitive.ObjectIDFromHex(currentUser.ID.Hex())

// 	// Apply filter and update
// 	filter := bson.D{{Key: "_id", Value: id}}
// 	update := bson.D{{Key: "$set", Value: bson.D{{Key: "commits", Value: currentUser.Commits}}}}
// 	res, err := usersCollection.UpdateOne(context.TODO(), filter, update)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	return res
// }

// // Returns true if user exists, false otherwise
// func UserExists(email string) bool {
// 	cursor, err := usersCollection.Find(context.TODO(), bson.D{{
// 		Key:   "email",
// 		Value: email,
// 	}})
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	// convert the cursor result to bson
// 	var results []bson.M
// 	if err = cursor.All(context.TODO(), &results); err != nil {
// 		log.Fatalf("Some error occured. Err: %s", err)
// 	}

// 	return len(results) != 0
// }

// // Returns a list of users in bson format
// func GetUsers() []bson.M {
// 	// retrieve all the documents in users collection
// 	cursor, err := usersCollection.Find(context.TODO(), bson.D{})
// 	if err != nil {
// 		log.Fatalf("Some error occured. Err: %s", err)
// 	}

// 	// convert the cursor result to bson
// 	var results []bson.M
// 	if err = cursor.All(context.TODO(), &results); err != nil {
// 		log.Fatalf("Some error occured. Err: %s", err)
// 	}

// 	return results
// }
