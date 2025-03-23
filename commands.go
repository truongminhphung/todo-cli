package main

import (
	"fmt"
	"sort"
)

// AddTask adds a new task with a unique ID
func AddTask(tasks []Task, description string) ([]Task, error) {
	// Generate a new ID for the task
	newID := len(tasks) + 1
	newTask := Task{
		ID:          newID,
		Description: description,
		Completed:   false,
	}

	return append(tasks, newTask), nil
}

// DeleteTask removes a task by its ID
func DeleteTask(tasks []Task, id int) ([]Task, error) {
	for i, task := range tasks {
		if task.ID == id {
			// Remove the task from the slice
			return append(tasks[:i], tasks[i+1:]...), nil
		}
	}
	return tasks, nil // Return the original slice if the task was not found
}

// MarkTaskCompleted marks a task as completed by its ID
func MarkTaskCompleted(tasks []Task, id int) ([]Task, error) {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = true
			return tasks, nil
		}
	}
	return tasks, nil // Return the original slice if the task was not found
}

// ListTasks returns all tasks
func ListTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	// sort tasks by ID for consistent display
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].ID < tasks[j].ID
	})

	// Print each task
	for _, task := range tasks {
		status := " "
		if task.Completed {
			status = "X"
		}
		fmt.Printf("%d: %s [%s]\n", task.ID, task.Description, status)
	}
}
