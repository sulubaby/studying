package piscine

func HashCode(dec string) string {
	result := ""
	length := len(dec)

	for _, c := range dec {
		hash := (int(c) + length) % 127
		if hash < 33 {
			hash += 33
		}
		result += string(rune(hash))
	}
	return result
}
