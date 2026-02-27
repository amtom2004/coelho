package main

import "fmt"

var jwtToken = 3000000 // private global variable - first letter small

const DockerHubToken = "hghjjkhjk" // public global variable - first letter small

func main() {
	var username string = "Aron"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T\n", username)

	var isSigned bool = false
	fmt.Println(isSigned)
	fmt.Printf("The variable is of type: %T\n", isSigned)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("The variable is of type: %T\n", smallValue)

	var smallFloat float32 = 255.45557767987978
	fmt.Println(smallFloat)
	fmt.Printf("The variable is of type: %T\n", smallFloat)

	// default values and aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("The variable is of type: %T\n", anotherVariable)

	var stringVariable string
	fmt.Println(stringVariable)
	fmt.Printf("The variable is of type: %T\n", stringVariable)

	// implicit type variable declaration and assignment, usable to declare variables outside methods (global)

	var numberOfUsers = 3
	fmt.Println(numberOfUsers)

	// no var style, using walrus symbol (:=), usable only inside a function

	website := "aronmathewtom.me"
	fmt.Println(website)

	// checking global variables

	fmt.Printf("The varible is of type: %T\n", jwtToken)
	fmt.Printf("The varible is of type: %T\n", DockerHubToken)
}
