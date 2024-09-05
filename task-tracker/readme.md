# Task Tracker

The [Task Tracker](https://roadmap.sh/projects/task-tracker) is a simple task-tracking tool from your terminal.

Idea from [roadmap.sh](https://roadmap.sh/golang/projects)

Project page URL: https://github.com/hung1299/go-projects/tree/main/task-tracker
## How to run

You need to clone this project into your local machine first.

by running these commands:

```bash
git clone https://github.com/hung1299/go-projects.git
cd task-tracker
```
Run the following command to build and run the project:

```bash
make build
or 
make build-window # for window

# To see the list of available commands
./task-tracker --help 

# Add a new task 
./task-tracker add <description>

# Modify an existing task 
./task-tracker update <ID> <description>

# Delete a task 
./task-tracker delete <ID>

# List all tasks
./task-tracker list
./task-tracker list todo
./task-tracker list in-progress
./task-tracker list done

# Change the task's status
./task-tracker mark-todo <ID>
./task-tracker mark-in-progress <ID>
./task-tracker mark-done <ID>
```
