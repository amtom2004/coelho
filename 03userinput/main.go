package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating: ")

	// comm ok or error ok, no try-catch; just getting the error as a variable simply

	input, _ := reader.ReadString('\n') // keep reading until newline comes
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("The rating variable is of type: %T\n", input)
}
