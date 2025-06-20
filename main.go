package main

import (
	"fmt"
	"os"
	"os/user"

	"monk/repl"
)

func main() {
	args := os.Args[1:]

	if len(args) > 1 {
		panic("cannot take more than 1 argument")
	}

	if len(args) == 1 {
		repl.Run(args[0], os.Stdout)
	} else {
		user, err := user.Current()

		if err != nil {
			panic(err)
		}

		fmt.Printf("Hello %s! This is the Monkey Programming Language!\n", user.Username)
		fmt.Printf("Feel free to type in commands\n")
		repl.Start(os.Stdin, os.Stdout)
	}

}
