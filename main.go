package main

import (
	// "encoding/json"
	"fmt"
	"bufio"
	"os"
)

func main() {
	for {
		fmt.Print("\033[H\033[2J")
		choice := DisplayMenu()
		HandleMenuOption(choice)
		fmt.Println("Press Enter to continue...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}
