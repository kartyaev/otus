package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	length := len(s)
	if length == 0 {
		return "", nil
	}
	if unicode.IsDigit(rune(s[0])) {
		return "", ErrInvalidString
	}
	asIsNext := false
	var result = strings.Builder{}
	for i := 0; i < length; i++ {
		if s[i] == '\\' && !asIsNext {
			asIsNext = true
			continue
		}
		if asIsNext && !unicode.IsDigit(rune(s[i])) && s[i] != '\\' {
			return "", ErrInvalidString
		}
		if i+1 < length && unicode.IsDigit(rune(s[i+1])) {
			if i+2 < length && unicode.IsDigit(rune(s[i+2])) {
				return "", ErrInvalidString
			}
			counter, _ := strconv.Atoi(string(s[i+1]))
			result.WriteString(strings.Repeat(string(s[i]), counter))
			i++
			asIsNext = false
		} else {
			result.WriteByte(s[i])
			asIsNext = false
		}
	}
	return result.String(), nil
}
