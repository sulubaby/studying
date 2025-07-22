package piscine

func RepeatAlpha(s string) string {
	result := ""
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			for i := 0; i < int(c-'A'+1); i++ {
				result += string(c)
			}
		} else if c >= 'a' && c <= 'z' {
			for i := 0; i < int(c-'a'+1); i++ {
				result += string(c)
			}
		} else {
			result += string(c)
		}
	}
	return result
}
