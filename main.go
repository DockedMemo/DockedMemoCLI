package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var ARGS = os.Args

func main() {

	switch ARGS[1] {
	case "help":
		handleHelp(getArg(2))
	case "add":
		handleAdd(getArg(2), getArg(3))
	case "remove":
		handleRemove(getArg(2))
	case "list":
		handleList(getArg(2))
	}
}

func handleHelp(term string) {
	switch term {
		case "add":
			fmt.Println("add help")
		case "remove":
			fmt.Println("remove help")
		case "list":
			fmt.Println("list help")
		default:
			fmt.Println("general help")
	}
}

func handleAdd(name string, description string) {
	if (name == "") {
		log.Fatal("You must specify the name of the task to add")
	}
	
	AddTask(Task{name, description, time.Now().Unix()})
	fmt.Println("Successfully added task")
}

func handleRemove(name string) {
	if (name == "") {
		log.Fatal("You must specify the name of the task to remove")
	}

	RemoveTask(name)
	fmt.Println("Successfully removed task")
}

func handleList(filter string) {
	var outputMessage strings.Builder

	tasks := GetTasks()

	for i, task := range tasks {
		outputMessage.WriteString( fmt.Sprintf("%v. %v\n  > %v\n", i+1, task.Name, task.Description))
	}

	fmt.Println(outputMessage.String())
}

// Returns an empty string if the index is out of bounds
func getArg(index int) string {
	if (len(ARGS) <= index) {
		return ""
	}
	return ARGS[index]
}