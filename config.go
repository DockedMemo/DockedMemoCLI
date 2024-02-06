package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/kirsle/configdir"
)

type Task struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Created int64 `json:"created"`
}

func getConfigDir() string {
	configPath := configdir.LocalConfig("DockedMemoCLI")

	err := configdir.MakePath(configPath) // Ensure it exists.
	if err != nil {
		panic(err)
	}

	return configPath
}

func getConfigFile() []byte {
	data, err := os.ReadFile(filepath.Join(getConfigDir(), "tasks.json"))
	
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func writeToFile(tasks []Task) {
	// Encode our new list of tasks as JSON
	encodedString, err := json.Marshal(tasks)
	if err != nil {
		log.Fatal(err)
	}

	// Write the file to disk
	os.WriteFile(filepath.Join(getConfigDir(), "tasks.json"), encodedString, os.FileMode(0644))
}

func GetTasks() []Task {
	todoItems := []Task{}

	json.Unmarshal(getConfigFile(), &todoItems)

	return todoItems
}

func AddTask(task Task) {
	currentTasks := GetTasks()
	// Ensure that there are no tasks with the same name
	for _, currentTask := range currentTasks {
		if (currentTask.Name == task.Name) {
			log.Fatal("Cannot have two tasks with the same name")
		}
	}

	// Append the task
	currentTasks = append(currentTasks, task)
	
	writeToFile(currentTasks)
}

func RemoveTask(taskName string) {
	currentTasks := GetTasks()
	haveRemovedStuff := false

	for i, currentTask := range currentTasks {
		if (currentTask.Name == taskName) {
			currentTasks = remove(currentTasks, i)
			haveRemovedStuff = true
		}
	}

	if (!haveRemovedStuff) {
		log.Fatal("Could not remove the task, it does not exist in the *database*")
	}

	writeToFile(currentTasks)
}

// TODO: Use generics
func remove(activeSlice []Task, i int) []Task {
    activeSlice[i] = activeSlice[len(activeSlice)-1]
    return activeSlice[:len(activeSlice)-1]
}
