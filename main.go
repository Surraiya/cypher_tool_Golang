package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Cypher Tool!")
	fmt.Println("\nSelect operation (1/2):\n1. Encrypt.\n2. Decrypt.")

	// Get Input
	//Get the input as number so if any other input is given, it shows a invalid message
	operation := 0
	fmt.Print(fmt.Scanln(&operation))
	//If its empty then the message should be

	fmt.Println("\nSelect cypher (1/2):\n1. ROT13.\n2. Reverse.")

	//Get Input

	fmt.Println("\nEnter the message:")

	//Get Input

	fmt.Println("\nDecrypted message using reverse:")

	//return output
}

// Get the input data required for the operation
func getInput() (toEncrypt bool, encoding string, message string) {
	//Suppose toEncrypt = true
	//encoding = "ROT13"
	//message = "hello"
	//It means I need to encrypt "hello" message using the preferred encryption technique
	return
}
