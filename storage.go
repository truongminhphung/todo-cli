package main

import (
	"encoding/json"
	"os"
)

const TASK_FILE = "tasks.json"

func loadTasks() ([]Task, error) {
	file, err := os.Open(TASK_FILE)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil // Return an empty slice if the file doesn't exist
		}

		return nil, err // Return the error if there was a problem opening the file
	}
	defer file.Close()

	// Decode the JSON data from the file into a slice of Task structs
	var tasks []Task
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// Write task to the JSON file
func saveTasks(tasks []Task) error {
	// create or over write the file
	file, err := os.Create(TASK_FILE)
	if err != nil {
		return err
	}
	defer file.Close()
	// Encode the tasks slice to JSON and write it to the file
	encoder := json.NewEncoder(file)
	err = encoder.Encode(tasks)
	if err != nil {
		return err
	}
	return nil
}
