package main

import (
	"fmt"

	"example.com/note/note"
)

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")


	return title, content
}

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(userNote)

}

func getUserInput(prompt string) (string) {
	fmt.Print(prompt)
	var value string
	fmt.Scanln(&value)


	return value
}
