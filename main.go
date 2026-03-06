package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes/note"
	"example.com/notes/todo"
)

type saver interface {
	Save() error
}

type outputtable interface {
	Display()
	saver
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text:")

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userNote)

	if err != nil {
		return
	}

	outputData(todo)
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()
	dataTypeText := ""

	switch data.(type) {
	case note.Note:
		dataTypeText = "note"
	case todo.Todo:
		dataTypeText = "todo"
	}

	if err != nil {
		fmt.Printf("Saving %v failed.\n", dataTypeText)
		return err
	}

	fmt.Printf("Saving %v succeded!\n", dataTypeText)
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note Content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
