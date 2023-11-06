package main

import (
	"bufio"
	"cypher_tool/caesar"
	"cypher_tool/reverse"
	"cypher_tool/rot13"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		toEncrypt := false
		encoding := ""

		fmt.Println("Welcome to the Cypher Tool!")

		fmt.Println("\nSelect operation (1/2):\n1. Encrypt.\n2. Decrypt.")
		toEncrypt = toEncryption(ChoiceInput())

		fmt.Println("\nSelect cypher (1/3):\n1. ROT13.\n2. Reverse.\n3. Caesar")
		encoding = getEncryptionType(ChoiceInput())

		fmt.Println("Enter a message: ")
		message := getMessage()

		getInput(toEncrypt, encoding, message)

		if !askForAnotherOperation() {
			break
		}
	}
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
		return "Caesar"
	}
}

func ChoiceInput() int {
	var choice int
	for {
		if _, err := fmt.Scan(&choice); err == nil {
			if choice == 1 || choice == 2 || choice == 3 {
				break
			}
		}
		fmt.Println("Invalid Operation. Please Enter Again:")
	}
	return choice
}

func getMessage() string {
	reader := bufio.NewReader(os.Stdin)

	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input != "" {
			return input
		}
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
	case "Caesar":
		var shift int
		fmt.Println("Enter the Caesar Cipher shift you want: ")
		_, err := fmt.Scan(&shift)
		if err != nil {
			fmt.Println("Error reading the shift:", err)
			return
		}

		if toEncrypt {
			fmt.Printf("Encrypted message using %s technique:\n%s\n", encoding, caesar.Encrypt_caesar(message, shift))
		} else {
			fmt.Printf("Decrypted message using %s technique:\n%s\n", encoding, caesar.Decrypt_caesar(message, shift))
		}
	}
}

func askForAnotherOperation() bool {
	for {
		fmt.Println("\nDo you want to perform another operation? (y/n)")
		var anotherOperation string
		_, err := fmt.Scan(&anotherOperation)
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		if anotherOperation != "y" && anotherOperation != "n" {
			fmt.Println("Invalid operation. Please enter 'y' or 'n'.")
			continue
		}
		return anotherOperation == "y"
	}
}
