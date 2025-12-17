# Aslsh

A simple shell written in Go for learning purposes.

## Features

- Customizable prompt via `.aslshrc`
- Basic built-in commands: `echo`, `pwd`, `clear`, `history`, `help`, `exit`, `cd`, `ls`, `cat`, `touch`, `rm`, `cp`, `mv`
- Command history saved to `.history`
- Modular code structure for easy extension

## Built-in Commands

| Command   | Description                                 |
|-----------|---------------------------------------------|
| `echo`    | Print arguments to the terminal             |
| `pwd`     | Print current working directory             |
| `clear`   | Clear the terminal screen                   |
| `history` | Show command history                        |
| `help`    | Show help for built-in commands             |
| `exit`    | Exit the shell                              |
| `cd`      | Change the current directory                |
| `ls`      | List files and directories                  |
| `cat`     | Display contents of a file                  |
| `touch`   | Create an empty file                        |
| `rm`      | Remove a file                               |
| `cp`      | Copy files                                  |
| `mv`      | Move or rename files                        |

## Getting Started

### Prerequisites

- Go 1.24 or newer

### Build

```sh
cd src
go mod tidy   # Ensure dependencies are installed
go build -o aslsh aslsh.go
```

### Run

```sh
./aslsh
```

## Configuration

You can customize the shell prompt by editing the `.aslshrc` file in the `src` directory:

```
ps=your_prompt_here
```

If `.aslshrc` does not exist, the shell will use a default prompt.

## Project Structure

- `src/aslsh.go` - Main shell implementation
- `src/commands/` - Built-in command implementations
- `src/util/` - Utility functions (user input, history, etc.)
- `.aslshrc` - Shell configuration file (in `src/`)
- `.history` - Command history file (in `src/`)

## Testing

Unit tests are provided for built-in commands. To run the tests:

```sh
cd src
go test -v ./...
```

Test results will be shown in the terminal, indicating which tests passed or failed.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests for improvements or new features.

## License

This project is for educational purposes.