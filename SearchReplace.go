package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	str := os.Args[1]
	from := rune(os.Args[2][0])
	to := rune(os.Args[3][0])

	for _, c := range str {
		if c == from {
			z01.PrintRune(to)
		} else {
			z01.PrintRune(c)
		}
	}
	z01.PrintRune('\n')
}
