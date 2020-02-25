package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/ayushbhargav/monkey-lang/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Welcome %s to Monkey programming language. V0.0.1\n", user.Username)
	fmt.Printf("Type in commands to continue\n")
	repl.Start(os.Stdin, os.Stdout)
}
