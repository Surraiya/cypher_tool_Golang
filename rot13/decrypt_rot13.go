func decrypt_rot13(s string) string {
	return encrypt_rot13(s)
}
//since rot13 is symmetric encrpytion we can just decrypt using the encryption again.