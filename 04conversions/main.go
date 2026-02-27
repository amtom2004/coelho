package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter your rating: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n') // Reads value as "4/n"
	fmt.Println("You've rated as ", input)
	newRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // strconv.ParseFloat("4/n", 64) fails, so strings.Trimspace() removes whitespace characters
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The modified rating is ", newRating+1)
	}
}
