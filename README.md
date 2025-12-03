# Aslsh

A simple shell written in Go for learning purposes.

## Features

- Customizable prompt via `.aslshrc`
- Basic built-in commands: `echo`, `pwd`, `clear`, `history`, `help`, `exit`
- Command history saved to `.history`
- Modular code structure for easy extension

## Getting Started

### Prerequisites

- Go 1.24 or newer

### Build

```sh
cd src
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

## Project Structure

- `src/aslsh.go` - Main shell implementation
- `src/commands/` - Built-in command implementations
- `src/util/` - Utility functions (user input, history, etc.)
- `.aslshrc` - Shell configuration file
- `.history` - Command history file

## Testing

Unit tests are provided for built-in commands. To run the tests:

```sh
cd src
go test -v ./testing
```

You can also run all tests in the project (if you add more):

```sh
go test -v ./...
```

The tests cover command behaviors such as `cd` and `pwd`.  
Test results will be shown in the terminal, indicating which tests passed or failed.

## License

This project is for educational purposes.