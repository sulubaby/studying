package piscine

func RepeatAlpha(s string) string {
	result := ""
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			for i := 'A'; i < i <= c; i++ {
				result += string(c)
			}
		} else if c >= 'a' && c <= 'z' {
			for i := 'a'; i <= c; i++ {
				result += string(c)
			}
		} else {
			result += string(c)
		}
	}
	return result
}
