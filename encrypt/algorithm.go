package encrypt

func EnCoup(plainText string) string {
	var encryptedText string
	for i := 0; i < len(plainText); i++ {
		encryptedText = encryptedText + string(plainText[i]+3)
	}
	return encryptedText
}
