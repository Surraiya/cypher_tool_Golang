package main

import (
	"fmt"
	"strings"
)

func main() {
	toEncrypt := false
	encoding := ""
	message := ""

	fmt.Println("Welcome to the Cypher Tool!")

	fmt.Println("\nSelect operation (1/2):\n1. Encrypt.\n2. Decrypt.")
	toEncrypt = toEncryption(ChoiceInput())
	fmt.Println(toEncrypt)

	fmt.Println("\nSelect cypher (1/2):\n1. ROT13.\n2. Reverse.")
	encoding = getEncryptionType(ChoiceInput())
	fmt.Println(encoding)

	fmt.Println("\nEnter the message:")
	message = getMessage()
	fmt.Println(message)

	fmt.Println("\nDecrypted message using reverse:")
	getInput(toEncrypt, encoding, message)
}

func getInput(toEncrypt bool, encoding string, message string) {
	//Suppose toEncrypt = true
	//encoding = "ROT13"
	//message = "hello"
	//It means I need to encrypt "hello" message using the preferred encryption technique
	if toEncrypt && encoding == "ROT13" {
		fmt.Println(rot13.encrypt_rot13(message))
	} else {
		fmt.Println(rot13.decrypt_rot13(message))
	}

	if toEncrypt && encoding == "Reverse" {
		fmt.Println(reverse.encrypt_reverse(message))
	} else {
		fmt.Println(reverse.decrypt_reverse(message))
	}

	return
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

func getMessage() string {
	var input string

	for {
		_, _ = fmt.Scanln(&input)
		//remove whitespaces from the beginning and the end of the input
		trimmedInput := strings.TrimSpace(input)
		if trimmedInput != "" {
			return trimmedInput
		}
		fmt.Println("Invalid message. Please enter a non-empty message:")
	}
}

func toEncryption(operation int) bool {
	return operation == 1
}

func getEncryptionType(encryptionType int) string {
	if encryptionType == 1 {
		return "ROT13"
	} else if encryptionType == 2 {
		return "Reverse"
	} else {
		return "something else"
	}
}
