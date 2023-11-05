package rot13

func Encrypt_rot13(s string) string {

	shiftstring := ""
	for _, char := range s {

		if char >= 'a' && char <= 'z' {
			shift := int(char) - 'a' + 13
			shift = shift % 26
			shiftstring += string(rune('a' + shift))

		} else if char >= 'A' && char <= 'Z' {
			shift := int(char) - 'A' + 13
			shift = shift % 26
			shiftstring += string(rune('A' + shift))

		} else {
			shiftstring += string(char)
		}
	}

	return shiftstring
}

func Decrypt_rot13(s string) string {
	return Encrypt_rot13(s)
}

//since rot13 is symmetric encrpytion we can just decrypt using the encryption again.
