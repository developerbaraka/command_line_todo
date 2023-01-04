package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Todo struct {
	content string
	status  bool
}

type TodoList struct {
	items []Todo
}

func NewTodoList() *TodoList {
	return &TodoList{}
}

type Option struct {
	key     string
	message string
}

func main() {
	// DisplayMenu()

	todo := NewTodoList()

	for {
		// menu display
		DisplayMenu()

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("There was a problem in reading the input!")
		}

		input = strings.TrimSpace(input)
		args := strings.Fields(input)

		// check for input
		if len(args) < 1 {
			continue
		}

		switch args[0] {

		case "add":
			todo.Add()
		case "list":
			todo.List()
		case "remove":
			todo.Remove()
		case "done":
			todo.Done()
		case "exit":
			fmt.Println("You have exited the program!!")
			return
		default:
			fmt.Println("Use the right operation to continue (;")
		}
	}

}

func (t *TodoList) Add() {
	// read user input
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("There was a problem in reading the input!")
	}

	input = strings.TrimSpace(input)

	// push to array
	t.items = append(t.items, Todo{
		content: input,
		status:  false,
	})
}

func (t *TodoList) List() {
	// print the todo list
	if len(t.items) == 0 {
		fmt.Printf("\n You have no items yet!! \n")

	} else {
		fmt.Printf("\n Below are your todo items. \n")

		for key, val := range t.items {
			fmt.Printf("%d. %s :status(%v) \n", key+1, val.content, val.status)
		}
		fmt.Println("")
	}

}

func (t *TodoList) Remove() {
	// read user input
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("There was a problem in reading the input!")
	}

	input = strings.TrimSpace(input)

	// convert to integer for splicing array
	index, _ := strconv.Atoi(input)
	if index < 1 || index > len(t.items) {
		fmt.Println("Error: task number is out of range")
	}
	index = index - 1

	t.items = append(t.items[:index], t.items[index+1:]...)
}
func (t *TodoList) Done() {
	// read user input
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("There was a problem in reading the input!")
	}

	input = strings.TrimSpace(input)

	// convert to integer for splicing array
	index, _ := strconv.Atoi(input)
	if index < 1 || index > len(t.items) {
		fmt.Println("Error: task number is out of range")
	}
	index = index - 1

	item := &t.items[index]
	item.status = true

}

func DisplayMenu() {
	time.Sleep(time.Second)
	bold := color.New(color.FgWhite)
	white := bold.Add(color.Bold)
	white.Println("This is the menu you are going to use with this application!")

	color.White("--------- STARTING TODO APP ----------")
	fmt.Println("Use the key words to perform operations")

	options := []Option{
		{
			key:     "add",
			message: "Add an item to your todo list.",
		},
		{
			key:     "remove",
			message: "Remove an item by using its index displayed besides it.",
		},
		{
			key:     "list",
			message: "List all items in your todo list.",
		},
		{
			key:     "done",
			message: "Mark an item done using their index.",
		},
		{
			key:     "exit",
			message: "Exit the app.",
		},
	}

	for _, val := range options {
		color.Yellow("[X] %s \t :explanation(%v) \n", val.key, val.message)
	}

}
