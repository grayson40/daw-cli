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

1. To Start, clone the Daw repository from GitHub, change into the Daw directory, and checkout the master branch.
    ```
    git clone https://github.com/grayson40/daw-version-control.git
    cd daw-version-control
    git checkout master
    ```
    Working off the master branch will ensure that you're using the latest released version of Daw.

2. Build and install the module.
    ```
    go build
    go install
    ```
    
3. You have now successfully installed the daw version control tool.

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
