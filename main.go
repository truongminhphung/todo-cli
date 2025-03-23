package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: todo <command> [<args>]")
		fmt.Println("Commands: add <description>, delete <id>, complete <id>, list")
		os.Exit(1)
	}
	command := os.Args[1]

	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		os.Exit(1)
	}

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo add <description>")
			os.Exit(1)
		}
		description := os.Args[2]
		// var err error // Declare err to avoid shadowing
		tasks, err = AddTask(tasks, description) // Remove the redeclaration
		if err != nil {
			fmt.Println("Error adding task:", err)
			os.Exit(1)
		}
		fmt.Println("Task added:", tasks)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo delete <id>")
			os.Exit(1)
		}
		idStr := os.Args[2]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("Invalid ID:", err)
			os.Exit(1)
		}
		originalLength := len(tasks)
		tasks, err = DeleteTask(tasks, id)
		if err != nil {
			fmt.Println("Error deleting task:", err)
			os.Exit(1)
		}
		if len(tasks) < originalLength {
			fmt.Println("Task deleted:", id)
		} else {
			fmt.Println("Task not found:", id)
		}
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo complete <id>")
			os.Exit(1)
		}
		idStr := os.Args[2]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("Invalid ID:", err)
			os.Exit(1)
		}
		_, err = MarkTaskCompleted(tasks, id)
		if err != nil {
			fmt.Println("Error marking task as completed:", err)
			os.Exit(1)
		}
		fmt.Println("Task marked as completed:", id)
	case "list":
		ListTasks(tasks)
	default:
		fmt.Println("Unknown command:", command)
		fmt.Println("Commands: add <description>, delete <id>, complete <id>, list")
		os.Exit(1)
	}

	err = saveTasks(tasks)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		os.Exit(1)
	}
}
