package caesar

import "bytes"

func Encrypt_caesar(plaintext string, shift int) string {
	var result bytes.Buffer
	shift = shift % 26

	for _, char := range plaintext {
		if ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') {
			asciiOffset := 'a'
			if 'A' <= char && char <= 'Z' {
				asciiOffset = 'A'
			}
			result.WriteRune((char-asciiOffset+rune(shift))%26 + asciiOffset)
		} else {
			result.WriteRune(char) // Non-alphabet characters are left unchanged
		}
	}
	return result.String()
}

func Decrypt_caesar(ciphertext string, shift int) string {
	return Encrypt_caesar(ciphertext, 26-shift) // Decrypt by shifting in reverse
}
