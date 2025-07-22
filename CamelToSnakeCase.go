package piscine

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}

	result := ""

	for i, c := range s {
		if c >= 'A' && c <= 'Z' {
			if i > 0 {
				result += "_"
			}
			result += string(c + 32) // convert to lowercase
		} else {
			result += string(c)
		}
	}
	return result
}
