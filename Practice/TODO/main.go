package main

import (
	"bufio"
	"example.com/todo/note"
	"example.com/todo/todo"
	"fmt"
	"os"
	"strings"
)

type saver interface {
	Save() error
}

// embeded interface using interface inside interface
type outputtable interface {
	saver
	Display()
}

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
	err = outputData(todoItem)
	if err != nil {
		return
	}
	outputData(userNote)

}

//main function end

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}
func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("error while saving file")
		return err
	}
	fmt.Println("file saved")
	return nil
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
