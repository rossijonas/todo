<h1 align="center">todo</h1>

*<p align="center">A Simple Todo List for the CLI.</p>*

## About

Use `todo` CLI tool to create and manage a simple todo list.

### Status:

[![Actions Status](https://github.com/rossijonas/todo/workflows/Test/badge.svg)](https://github.com/rossijonas/todo/actions)
[![Actions Status](https://github.com/rossijonas/todo/workflows/Build/badge.svg)](https://github.com/rossijonas/todo/actions)

### Features:

- Cross platform:  Linux / Macos / Windows.

- Creates and manages a todo list.

- Support custom task list filename via `TODO_FILENAME` environment variable (`.json` format).

- Support adding tasks from STDIN.

_(This is an exercise from the book "Powerful Command-Line Applications in Go".)_

## Installation

### Requirements:

- [Go](https://go.dev/) version 1.18.6 (or above)

### How to install:

- Run: 

  ```
  $ go install github.com/rossijonas/todo/cmd/todo@latest
  ```

## Usage

### Options:

```
$ todo -h
todo - A Simple Todo List for the CLI.
Usage:
  -add
        Task to be included in the ToDo list
  -complete int
        Item to be completed
  -del int
        Item to be deleted
  -list
        List all tasks
```

### Examples:

#### Create a custom name for the todo list file:

```
$ TODO_FILENAME=MyTodoList.json
```

#### List all tasks:

```
$ todo -list
  1: My task 1
✔ 2: My task 2
  3: My task 3
  4: My new task
✔ 5: old task
```

#### Add a new task:

```
$ todo -add Wash my feet
```

#### Mark task as competed:

```
$ todo -complete 3
```

#### Delete task:

```
$ todo -del 2
```

## Backlog

- Add cover image to README file.

- Update the custom usage function to include additional instructions on how to provide new tasks to the tool.

- Add support for a `-v` (verbose) flag, displaying date of creation and date of completion.

- Add support for a `-pending` flag, displaying only pending tasks (hiding completed tasks).

- Add support for multiline input from STDIN in the `-add` flag, adding each line as a new task into the list.

