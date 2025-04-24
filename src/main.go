package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/psidh/Ganges/src/repl"
)

func main() {
	user, err := user.Current()

	if err != nil {
		fmt.Printf("Error : ", err)
		panic(err)
	}

	fmt.Println("This is Ganges Programming Language! \n", user.Username)

	fmt.Printf("Type your code...\n")
	repl.Start(os.Stdin, os.Stdout)

}
