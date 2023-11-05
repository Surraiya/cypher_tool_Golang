package main

import (
	"bufio"
	"cypher_tool/reverse"
	"cypher_tool/rot13"
	"fmt"
	"os"
	"strings"
)

func main() {
	toEncrypt := false
	encoding := ""

	fmt.Println("Welcome to the Cypher Tool!")

	fmt.Println("\nSelect operation (1/2):\n1. Encrypt.\n2. Decrypt.")
	toEncrypt = toEncryption(ChoiceInput())
	fmt.Println(toEncrypt)

	fmt.Println("\nSelect cypher (1/2):\n1. ROT13.\n2. Reverse.")
	encoding = getEncryptionType(ChoiceInput())
	fmt.Println(encoding)

	fmt.Println("Enter the message: ")
	message := getMessage()
	fmt.Printf("You entered the following message: %s\n", message)

	getInput(toEncrypt, encoding, message)
}

func getMessage() string {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter a message: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input != "" {
			return input
		}
		fmt.Println("Invalid input. Please try again.")
	}
}

func getInput(toEncrypt bool, encoding string, message string) {
	switch encoding {
	case "ROT13":
		if toEncrypt {
			fmt.Printf("Encrypted message using %s technique:\n%s\n", encoding, rot13.Encrypt_rot13(message))
		} else {
			fmt.Printf("Decrypted message using %s technique:\n%s\n", encoding, rot13.Decrypt_rot13(message))
		}
	case "Reverse":
		if toEncrypt {
			fmt.Printf("Encrypted message using %s technique:\n%s\n", encoding, reverse.Encrypt_reverse(message))
		} else {
			fmt.Printf("Decrypted message using %s technique:\n%s\n", encoding, reverse.Decrypt_reverse(message))
		}
	}
}

func ChoiceInput() int {
	var choice int
	for {
		if _, err := fmt.Scan(&choice); err == nil {
			if choice == 1 || choice == 2 {
				break
			}
		}
		fmt.Println("Invalid Operation. Please Enter Again:")
	}
	return choice
}

func toEncryption(operation int) bool {
	return operation == 1
}

func getEncryptionType(encryptionType int) string {
	switch encryptionType {
	case 1:
		return "ROT13"
	case 2:
		return "Reverse"
	default:
		return "something else"
	}
}
