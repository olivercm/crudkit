package utils

import (
	"strings"
	"unicode"
)

func LowerCaseFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

func UpperCaseFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func UpperCamel(name string) string {
	sb := strings.Builder{}
	upper := true
	for i := 0; i < len(name); i++ {
		c := name[i]
		if c == '_' {
			upper = true
			continue
		}
		if upper {
			if 'a' <= c && c <= 'z' {
				c = c + 'A' - 'a'
			}
			upper = false
		}
		sb.WriteByte(c)
	}
	return sb.String()
}






