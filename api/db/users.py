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

# Creates project dictionary
Project = lambda name, path, saved, changes: {
    "name": name,
    "path": path,
    "saved": saved,
    "changes": changes,
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
    if cursor == None:
        return None
    projects = []
    for project in cursor["projects"]:
        project_dict = Project(
            project["name"],
            project["path"],
            project["saved"],
            project["changes"],
        )
        projects.append(project_dict)
    return projects


# Delete all users for testing
def delete_all_users():
    cursor = users_collection.find()
    users = list(cursor)
    for user in users:
        user_id = user["_id"]
        users_collection.delete_one({"_id": ObjectId(user_id)})


# Returns true if project exists in db
def project_exists(user_projects, in_project):
    for project in user_projects:
        if project["path"] == in_project["path"]:
            return True
    return False


# Add project in db
def add_project(project, user_id):
    user_projects = get_user_projects(user_id)
    # If user has no projects
    if user_projects == None:
        users_collection.update_one(
            {"_id": ObjectId(user_id)},
            {
                "$set": {"projects": project},
            },
        )
    # Append to existing projects list
    else:
        if not project_exists(user_projects, project):
            user_projects.append(project)
        users_collection.update_one(
            {"_id": ObjectId(user_id)},
            {
                "$set": {"projects": user_projects},
            },
        )


# Update project changes
def update_project_changes(project_name, changes, user_id):
    users_collection.update_one(
        {
            "_id": ObjectId(user_id),
            "projects.name": project_name,
        },
        {"$set": {"projects.$.changes": changes}},
    )
