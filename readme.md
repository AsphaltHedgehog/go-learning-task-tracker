The Task Tracker

Check result of my work on the challenge from [Roadmap.sh](https://roadmap.sh/projects/task-tracker)\n

!!! Important !!!\n
On linux or windows app creates temp folder with json file for storing the tasks. After using the program u need manually remove folder with files\n

Folder created on Win and Linux in different places.\n
Win: %user%\AppData\Local\task-tracker-go\n
Easy to remove using win+r paste %localappdata%\n
Linux: /home/<user>/.local/share/go-task-tracker\n
!!! Important !!!\n

How to use:\n
Clone repo\n
git clone https://github.com/AsphaltHedgehog/go-learning-task-tracker\n

Run go build command\n
go build -o task-tracker-go\n
Run task-tracker-go without arguments to see all options in help section\n

A Task consists of only five fields:\n
ID int\n
Description string\n
Status (todo/in-progress/done)\n
UpdatedAt string\n
CreatedAt string\n

1. Add task by (specify only it's description)\n
   ./task-tracker-go add "Test run"\n

2. Update task by\n
   ./task-tracker-go update 5 "Take vacation and use it for hiking"\n

3. Delete taks by\n
   ./task-tracker-go delete 1\n

4. Update/mark task status by\n
   ./task-tracker-go mark todo 1\n
   ./task-tracker-go mark done 1\n
   ./task-tracker-go mark in-progress 5\n

5. List all task by\n
   ./task-tracker-go list done\n
   ./task-tracker-go list\n
   ./task-tracker-go list todo\n
   ./task-tracker-go list in-progress\n
