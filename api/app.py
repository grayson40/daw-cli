from flask import Flask, jsonify, request
from .db.users import *
from json import loads
from bson.json_util import dumps


app = Flask(__name__)

# User methods
@app.route("/user", methods=["GET", "POST"])
def user():
    # Get user from db
    if request.method == "GET":
        # Parse args
        user_email = request.args.get("email")
        user_name = request.args.get("username")
        user_id = request.args.get("id")

        # Get user by email or username
        if user_email != None:
            return get_user_by_email(user_email), "200"
        elif user_name != None:
            return get_user_by_user_name(user_name), "200"
        elif user_id != None:
            return get_user_by_id(user_id), "200"

    # Create user & post to db
    elif request.method == "POST":
        content_type = request.headers.get("Content-Type")
        if content_type == "application/json":
            json = request.json
            user_id = create_user(json["username"], json["email"], json["projects"])
            return user_id
        else:
            return "Content-Type not supported!"


# Return all users
@app.route("/users", methods=["GET"])
def users():
    return get_users()


# !!! FOR TESTING ONLY !!!
# Delete all users
@app.route("/delete", methods=["DELETE"])
def delete():
    delete_all_users()
    return "", "200"


# Project methods
@app.route("/projects", methods=["GET", "POST", "PUT"])
def projects():
    user_id = request.args.get("user_id")
    # Get all user projects
    if request.method == "GET":
        if user_id == None:
            return "", "403"
        return get_user_projects(user_id)
    # Add user project
    elif request.method == "POST":
        if user_id == None:
            return "", "403"
        content_type = request.headers.get("Content-Type")
        if content_type == "application/json":
            json = request.json
            project = Project(
                json["name"],
                json["path"],
                json["saved"],
                json["changes"],
            )
            add_project(project, user_id)
            return "", "204"
        else:
            return "Content-Type not supported!"
    # Update project changes
    elif request.method == "PUT":
        project_name = request.args.get("project_name")
        if user_id == None or project_name == None:
            return "", "403"
        content_type = request.headers.get("Content-Type")
        if content_type == "application/json":
            json = request.json
            changes = json["changes"]
            saved_time = json["saved"]
            update_project_changes(project_name, changes, saved_time, user_id)
            return "", "204"


# @app.route("/channels")
# def get_channels():
#     return jsonify(channel_rack)


# @app.route("/channels/instruments", methods=["GET", "POST"])
# def instruments():
#     if request.method == "GET":
#         return jsonify(channel_rack["instruments"])
#     elif request.method == "POST":
#         channel_rack["instruments"].append(request.get_json())
#         return "", 204


# @app.route("/channels/samplers", methods=["GET", "POST"])
# def samplers():
#     if request.method == "GET":
#         return jsonify(channel_rack["samplers"])
#     elif request.method == "POST":
#         channel_rack["samplers"].append(request.get_json())
#         return "", 204
