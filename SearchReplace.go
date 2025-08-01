package main 

import (
    "os"
    "github.com/01-edu/z01"
) 

func main() {
    if len(os.Args) != 4 {
        return 
    }
    result := ""
    str := os.Args[1]
    from :=os.Args[2]
    to :=os.Args[3]

    fromRune := rune(from[0])
    for _,c:= range str {
        if c == fromRune {
            result += to
        }else if c == 'é'{
            result +="o"
        } else {
            result += string(c)
        }
        
    }
    for _,a := range result {
            z01.PrintRune(a)
        }
    z01.PrintRune('\n')
}
