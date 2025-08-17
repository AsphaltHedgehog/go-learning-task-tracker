The Task Tracker

Check result of my work on the challenge from [Roadmap.sh](https://roadmap.sh/projects/task-tracker)

!!! Important !!!
On linux or windows app creates temp folder with json file for storing the tasks. After using the program u need manually remove folder with files

Folder created on Win and Linux in different places.
Win: %user%\AppData\Local\task-tracker-go
Easy to remove using win+r paste %localappdata%
Linux: /home/<user>/.local/share/go-task-tracker
!!! Important !!!

How to use:
Clone repo
git clone https://github.com/AsphaltHedgehog/go-learning-task-tracker

Run go build command
go build -o task-tracker-go
Run task-tracker-go without arguments to see all options in help section

A Task consists of only five fields:
ID int
Description string
Status (todo/in-progress/done)
UpdatedAt string
CreatedAt string

1. Add task by (specify only it's description)
   ./task-tracker-go add "Test run"

2. Update task by
   ./task-tracker-go update 5 "Take vacation and use it for hiking"

3. Delete taks by
   ./task-tracker-go delete 1

4. Update/mark task status by
   ./task-tracker-go mark todo 1
   ./task-tracker-go mark done 1
   ./task-tracker-go mark in-progress 5

5. List all task by
   ./task-tracker-go list done
   ./task-tracker-go list
   ./task-tracker-go list todo
   ./task-tracker-go list in-progress
