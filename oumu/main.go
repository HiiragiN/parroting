package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to the Parrot Program!")
	fmt.Println("Type something, and I will repeat it back. Type 'exit' to quit.")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("You: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			break
		}

		// Remove newline character from input
		input = input[:len(input)-1]

		if input == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		fmt.Println("Parrot:", input)
	}
}
