package main

import (
	"bufio"
	"fmt"
	"os"
)

// Note represents a note in the notepad
type Note struct {
	Title string
	Body  string
}

// Notepad represents the in-memory notepad
type Notepad struct {
	Notes []Note
}

// NewNotepad creates a new instance of Notepad
func NewNotepad() *Notepad {
	notepad := Notepad{
		Notes: []Note{},
	}
	return &notepad
}

// AddNote adds a new note to the notepad
func (n *Notepad) AddNote() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter note title: ")
	title, _ := reader.ReadString('\n')

	fmt.Print("Enter note body: ")
	body, _ := reader.ReadString('\n')

	note := Note{
		Title: title,
		Body:  body,
	}

	n.Notes = append(n.Notes, note)
	fmt.Println("Note added successfully!")
}

// ViewNotes displays all the notes in the notepad
func (n *Notepad) ViewNotes() {
	if len(n.Notes) == 0 {
		fmt.Println("No notes found.")
		return
	}

	for _, note := range n.Notes {
		fmt.Println("Title:", note.Title)
		fmt.Println("Body:", note.Body)
		fmt.Println("-----------------------")
	}
}

// UpdateNote updates a note in the notepad
func (n *Notepad) UpdateNote() {
	if len(n.Notes) == 0 {
		fmt.Println("No notes found.")
		return
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the title of the note to update: ")
	title, _ := reader.ReadString('\n')

	for i, note := range n.Notes {
		if note.Title == title {
			fmt.Print("Enter the new title: ")
			newTitle, _ := reader.ReadString('\n')

			fmt.Print("Enter the new body: ")
			newBody, _ := reader.ReadString('\n')

			note.Title = newTitle
			note.Body = newBody
			n.Notes[i] = note

			fmt.Println("Note updated successfully!")
			return
		}
	}

	fmt.Println("Note not found.")
}

// DeleteNote deletes a note from the notepad
func (n *Notepad) DeleteNote() {
	if len(n.Notes) == 0 {
		fmt.Println("No notes found.")
		return
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the title of the note to delete: ")
	title, _ := reader.ReadString('\n')

	for i, note := range n.Notes {
		if note.Title == title {
			n.Notes = append(n.Notes[:i], n.Notes[i+1:]...)
			fmt.Println("Note deleted successfully!")
			return
		}
	}

	fmt.Println("Note not found.")
}

func main() {
	notepad := NewNotepad()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("---------- In-Memory Notepad ----------")
		fmt.Println("1. Add Note")
		fmt.Println("2. View Notes")
		fmt.Println("3. Update Note")
		fmt.Println("4. Delete Note")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		choice, _ := reader.ReadString('\n')

		switch choice {
		case "1\n":
			notepad.AddNote()
		case "2\n":
			notepad.ViewNotes()
		case "3\n":
			notepad.UpdateNote()
		case "4\n":
			notepad.DeleteNote()
		case "5\n":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}

		fmt.Println()
	}
}
