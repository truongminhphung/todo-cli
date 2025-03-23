package main

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
