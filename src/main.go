package main

import (
	"fmt"
	"os"
	"os/user"
	"waiig/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type any commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
