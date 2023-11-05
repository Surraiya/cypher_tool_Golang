package reverse

func Encrypt_reverse(s string) string {
	var reversed []rune
	for _, char := range s {
		if ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') {
			if 'a' <= char && char <= 'z' {
				reversed = append(reversed, 'z'-(char-'a'))
			} else {
				reversed = append(reversed, 'Z'-(char-'A'))
			}
		} else {
			reversed = append(reversed, char) // Non-alphabet characters are left unchanged
		}
	}

	return string(reversed)
}

func Decrypt_reverse(s string) string {
	return Encrypt_reverse(s)
}
