package main

import (
	"bufio"
	"example.com/todo/note"
	"example.com/todo/todo"
	"fmt"
	"os"
	"strings"
)

func main() {

	title, content := getNoteData()

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	todoText := getUserInput("Enter todo: ")

	todoItem, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	todoItem.Display()

	err = todoItem.Save()
	if err != nil {
		fmt.Println("error while saving todo file")
		return
	}
	fmt.Println("todo file saved")

	userNote.Display()

	err = userNote.Save()
	if err != nil {
		fmt.Println("error while saving file")
		return
	}
	fmt.Println("file saved")
}

func getNoteData() (string, string) {
	title := getUserInput("Enter title: ")
	content := getUserInput("Enter the content: ")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	text = strings.TrimSpace(text)
	return text
}
