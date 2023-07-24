package decrypt

func DeCoup(encryptedText string) string {
	var decryptedText string
	for i := 0; i < len(encryptedText); i++ {
		decryptedText = decryptedText + string(encryptedText[i]-3)
	}
	return decryptedText
}
