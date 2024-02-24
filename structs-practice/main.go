package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

type outputtable interface {
	saver
	Display()
}

// type outputtable interface {
// 	Save() error
// 	Display()
// }

func getTodoData() string {
	return getUserInput("Todo text:")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")


	return title, content
}

func main() {
	printSomething(1)
	title, content := getNoteData()
	todoText := getTodoData()

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(todo)

	if err != nil {
		return
	}
	userNote, err := note.New(title, content)
	err = outputData(userNote)
}

func printSomething(value any ) {
	intVal, ok := value.(int)

	if ok {
		fmt.Println(intVal + 1)
	}

	floatVal, ok := value.(float64)

	if ok {
		fmt.Println(floatVal + 1)
	}

	stringVal, ok := value.(string)

	if ok {
		fmt.Println(stringVal)
	}


	switch value.(type) {
	case string:
		fmt.Println("This is a string")
	case float64:
		fmt.Println("This is a float")
	case int: 
		fmt.Println("This is an int")
	}
	fmt.Println(value)
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)

}

func saveData(data saver) error {
	err:= data.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}
	fmt.Println("Saving the note succeeded!")
	return nil
}

func getUserInput(prompt string) (string) {
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
