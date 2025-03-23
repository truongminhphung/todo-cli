# Todo CLI Application

A simple command-line todo list application written in Go.

## Building the Application

To build the application, run:

```bash
go build
```

This will create an executable named `todo-cli` in your current directory.

## Using the Application

### Listing Tasks

To list all tasks:

```bash
./todo-cli list
```

### Adding Tasks

To add a new task:

```bash
./todo-cli add "Your task description here"
```

Examples:
```bash
./todo-cli add "Buy Milk"
./todo-cli add "Buy a bunch of bananas"
```

### Completing Tasks

To mark a task as complete:

```bash
./todo-cli complete <task_id>
```

Replace `<task_id>` with the ID of the task you want to mark as complete.

### Deleting Tasks

To delete a task:

```bash
./todo-cli delete <task_id>
```

Replace `<task_id>` with the ID of the task you want to delete.

## Example Workflow

1. Build the application
   ```bash
   go build
   ```

2. Add some tasks
   ```bash
   ./todo-cli add "Buy groceries"
   ./todo-cli add "Finish project report"
   ```

3. View your tasks
   ```bash
   ./todo-cli list
   ```

4. Complete a task (assuming the task ID is 1)
   ```bash
   ./todo-cli complete 1
   ```

5. Delete a task (assuming the task ID is 2)
   ```bash
   ./todo-cli delete 2
   ```

6. View your updated task list
   ```bash
   ./todo-cli list
   ```