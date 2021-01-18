package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/samitcheema/Monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Welcome to the Monkey programming language %s, enjoy your stay!\n", user.Username)

	fmt.Printf("You may start typing commands--\n")
	repl.Start(os.Stdin, os.Stdout)
}
