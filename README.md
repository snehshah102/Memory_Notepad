#Memory Notepad

The In-Memory Notepad is a simple Go project that allows you to create, view, update, and delete notes in an in-memory database. It provides a basic command-line interface (CLI) for interacting with the notepad.

## Features

- **Add Note**: Create a new note by providing a title and body.
- **View Notes**: Display all the notes in the notepad.
- **Update Note**: Update the title and body of an existing note.
- **Delete Note**: Remove a note from the notepad.

## Prerequisites

Make sure you have the following software installed on your system:

- Go version go1.20.5 

## Getting Started

To get started with the In-Memory Notepad:

1. Clone this repository to your local machine.
   ```
   git clone https://github.com/your-username/in-memory-notepad.git
   ```

2. Open a terminal and navigate to the project directory.

3. Build and run the application using the following command:
   ```
   go run main.go
   ```

4. Follow the prompts in the terminal to interact with the notepad.

## Dependencies

This project relies on the following third-party libraries:

- `bufio`: For reading user input from the command-line.
- `fmt`: For formatted printing and string manipulation.
- `os`: For accessing the operating system's standard I/O.
- `fyne`: For creating a basic user interface (UI).

The necessary Go packages will be automatically downloaded and installed when you build and run the application.

## Contributing

Contributions are welcome! If you have any ideas, suggestions, or bug reports, please open an issue or submit a pull request.
