package piscine

func LastWord(s string) string {
	start := -1
	end := -1

	// Step 1: Loop backward to find the end index of the last word
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' && end == -1 {
			end = i + 1
		}
		if s[i] != ' ' {
			start = i
		}
		if s[i] == ' ' && end != -1 {
			break
		}
	}

	if start == -1 || end == -1 {
		return "\n" // no word found
	}

	return s[start:end] + "\n"
}
