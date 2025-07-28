package piscine

func Firstword (s string) strimg {
	result := ""
	for _, v := range s {
		if v == ' '{
			if len(result) != 0 {
				break
			}
		} else {
			result += string(v)
		}
	}
	return result
}
