import pymongo
import os
from bson.json_util import dumps
from bson.objectid import ObjectId
from dotenv import load_dotenv

# Creates user dictionary
User = lambda username, email, projects: {
    "username": username,
    "email": email,
    "projects": projects,
}

# Load env variables
load_dotenv()

# Connect to mongodb client
client = pymongo.MongoClient(os.environ.get("MONGODB_URL"))
db = client.test

# Connect to db
db = client.daw

# Get users collection
users_collection = db.users

# Returns list of users in db
def get_users():
    # Get users
    cursor = users_collection.find()
    users_cursor = list(cursor)

    # Converting to json
    users_list = dumps(users_cursor, indent=4)

    return users_list


# Post user to db
def create_user(username, email, projects):
    user = User(username=username, email=email, projects=projects)
    user_id = users_collection.insert_one(user).inserted_id
    return dumps({"_id": ObjectId(user_id)}, indent=4)


# Get user by email
def get_user_by_email(email):
    cursor = users_collection.find_one({"email": email})
    user = dumps(cursor, indent=4)
    return user


# Get user by username
def get_user_by_user_name(user_name):
    cursor = users_collection.find_one({"username": user_name})
    user = dumps(cursor, indent=4)
    return user


# Get user by id
def get_user_by_id(user_id):
    cursor = users_collection.find_one(ObjectId(user_id))
    user = dumps(cursor, indent=4)
    return user


# Get user projects
def get_user_projects(user_id):
    cursor = users_collection.find_one(ObjectId(user_id))
    user = dumps(cursor, indent=4)
    return user["projects"]


# Delete all users for testing
def delete_all_users():
    cursor = users_collection.find()
    users = list(cursor)
    for user in users:
        user_id = user["_id"]
        users_collection.delete_one({"_id": ObjectId(user_id)})
