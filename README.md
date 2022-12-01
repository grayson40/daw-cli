# DAW version control

DAW is a CLI client for DAW project file version control.

## Getting Started

Follow these instructions to build and install daw version control

## Pre-requisites

<details>
	<summary>Go</summary>

This tool is written and built with Golang. Download the latest version of Golang [here.](https://go.dev/doc/install)
</details>

<details>
	<summary>Git</summary>

Git is used to manage the codebase. Download the latest version of Git [here.](https://gitforwindows.org/)
</details>

## Installation & Running

**Getting started with daw version control...**

1. Click the green code button on the home page of this repository.
2. Copy the HTTPS URL provided.
3. Open a command prompt and navigate to your desired directory.
4. Clone the repository onto your local machine by executing the command:
    ```
    git clone [URL]
    ```
5. Navigate to the repository directory by executing the command:
    ```
    cd daw-version-control
    ```
6. Build and install the module by executing the commands:
    ```
    go build
    go install
    ```
7. You have now successfully installed the daw version control tool.

## Available Scripts

With the module installed, you can run:

### `add`

Add the project file(s) to be tracked.

**Example:**
```
daw add <file>...
```
### `commit`

Commit staged project file(s) with a specified message.

**Example:**
```
daw commit <message>
```
### `push`

Push the staged commits up the current ref.

**Example:**
```
daw push
```

### `status`

Displays paths that have differences between the index file and the current HEAD commit

**Example:**
```
daw status
```
