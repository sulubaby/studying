package piscine

func LastWord(s string) string {
	word := ""
	start := false

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			word = string(s[i]) + word
			start = true
		} else if start {
			break
		}
	}
	return word + "\n"
}
