package commands

var Msg = `cd       Change the current directory
exit     Exit the shell
pwd      Print the current working directory
echo     Print arguments to the terminal
export   Set environment variables
unset    Unset environment variables
alias    Define command shortcuts
help     Show help for built-in commands
history  Show command history
source   Reload .aslshrc
clear    Clear the terminal screen`

func Help() string {
    return Msg
}