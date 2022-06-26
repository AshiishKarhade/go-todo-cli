
# GO TODO CLI

A beautiful CLI To-do app made using Go!




## Installation and Run

Install go-cli-project with go

```bash
  git clone https://github.com/AshiishKarhade/go-todo-cli
  cd go-todo-cli 
```

Run go-cli-project with go

```bash
  go build ./cmd/todo 
  ./todo -add task-1
  ./todo -add task-2
  ./todo -list 
```
    
## Screenshots

![App Screenshot](https://github.com/AshiishKarhade/go-todo-cli/blob/main/cli-app.png)


## Usage/Examples

Adding task to todo list 
```bash
./todo -add task-1
./todo -add task-2
```
Marking a task as completed 
```bash
./todo -complete=1
./todo -complete=2
```
Deleting a task from todo list 
```bash
./todo -delete=1
./todo -delete=2
```
List all the tasks
```bash
./todo -list
```
