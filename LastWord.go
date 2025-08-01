package piscine

func LastWord(s string) string {
length := len(s)
if length == 0 {
    return "\n"
}

end := length - 1
for end >= 0 && s[end] == ' '{
    end --
}
if end < 0 {
    return "\n"
}
start := end
for start >= 0 && s[start] != ' '{
    start --
}
return s[start+1:end+1] + "\n"
}
