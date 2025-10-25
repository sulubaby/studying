package piscine

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}

	for i := 0; i < len(s); i++ {
		c := s[i]
		if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')) {
			return s 
		}
		if i == len(s)-1 && (c >= 'A' && c <= 'Z') {
			return s 
		}
		if i < len(s)-1 && (c >= 'A' && c <= 'Z') && (s[i+1] >= 'A' && s[i+1] <= 'Z') {
			return s 
		}
	}
	result := ""
	for i := 0; i < len(s); i++ {
		c := s[i]
		if (c >= 'A' && c <= 'Z') && i > 0 {
			result += "_"
		}
		result += string(c)
	}
	return result
}
