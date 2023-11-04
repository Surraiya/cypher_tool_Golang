func encrypt_rot13(s string) string {

	shiftstring := ""
	for _, char:= range s {
		
		if char >= 'a' && char <= 'z' {
			shift := int(char) - 'a' + 13
			shift = shift % 26
			shiftstring += string('a' + shift)
		
			} else if char >= 'A' && char <= 'Z' {
			shift := int(char) - 'A' + 13
			shift = shift % 26
			shiftstring += string('A' + shift)
		
			} else {
			shiftstring += string(char)
		}
	}
	
	return shiftstring
	}
	