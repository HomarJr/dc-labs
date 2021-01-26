package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Fatal error, not enough arguments")
		return
	} else {
		message := ""
		for i := 1; i < len(os.Args); i++ {
			message += os.Args[i] + " "
		}
		fmt.Println("Welcome to the jungle, " + message[:len(message)-1])
	}
}
