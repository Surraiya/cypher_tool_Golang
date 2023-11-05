func encryptreverse(s string) string {
	result := []rune{}
	for _, r := range s {
		result = append(result, ReverseAlphabet(r))
	}
	return string(result)
}

func ReverseAlphabet(ch rune) rune {
	if ch >= 'a' && ch <= 'z' {
		ch = 'z' - (ch - 'a')
		return ch
	} else if ch >= 'A' && ch <= 'Z' {
		ch = 'Z' - (ch - 'A')
		return ch
	} else if ch >= '0' && ch <= '9' {
		ch = '9' - (ch - '0')
		return ch
	}
	return ch
}
