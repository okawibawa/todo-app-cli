package main

import (
	"github.com/okawibawa/todo-app-cli/cmd"
)

func main() {
	cmd.Execute()

	// Uncomment the code below to register a signal channel to listen to command such as `ctrl + c` or `kill` before stopping the program.
	// sigchan := make(chan os.Signal, 1)
	// signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	//
	// <-sigchan
}
