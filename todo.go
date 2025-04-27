package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Todo struct {
	Description string
	Done        bool
}

var todos []Todo

const filename = "todos.json"

func loadTodos() {
	data, err := ioutil.ReadFile(filename)
	if err == nil {
		json.Unmarshal(data, &todos)
	}
}

func saveTodos() {
	data, _ := json.MarshalIndent(todos, "", "  ")
	ioutil.WriteFile(filename, data, 0644)
}

func addTask(desc string) {
	todos = append(todos, Todo{Description: desc})
	saveTodos()
}

func listTasks() {
	if len(todos) == 0 {
		fmt.Println("No tasks yet!")
		return
	}
	for i, t := range todos {
		status := "[ ]"
		if t.Done {
			status = "[x]"
		}
		fmt.Printf("%d. %s %s\n", i+1, status, t.Description)
	}
}

func markDone(index int) {
	if index > 0 && index <= len(todos) {
		todos[index-1].Done = true
		saveTodos()
	} else {
		fmt.Println("Invalid task number.")
	}
}

func main() {
	loadTodos()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- Todo List ---")
		fmt.Println("1. Add task")
		fmt.Println("2. List tasks")
		fmt.Println("3. Mark task as done")
		fmt.Println("4. Quit")
		fmt.Print("Choose an option: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Print("Enter task description: ")
			desc, _ := reader.ReadString('\n')
			desc = strings.TrimSpace(desc)
			addTask(desc)
			fmt.Println("Task added!")

		case "2":
			listTasks()

		case "3":
			listTasks()
			fmt.Print("Enter task number to mark done: ")
			numStr, _ := reader.ReadString('\n')
			numStr = strings.TrimSpace(numStr)
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("Invalid number.")
				continue
			}
			markDone(num)
			fmt.Println("Task marked as done!")

		case "4":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid option, please choose again.")
		}
	}
}
