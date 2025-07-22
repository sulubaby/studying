package piscine

func FirstWord(s string) string {
	word := ""
	for _, char := range s {
		if char == ' ' {
			break
		}
		word += string(char)
	}
	return word + "\n"
}
