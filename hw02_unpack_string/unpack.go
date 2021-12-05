package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	length := len([]rune(s))
	if length == 0 {
		return "", nil
	}
	if unicode.IsDigit(rune(s[0])) {
		return "", ErrInvalidString
	}
	asIsNext := false
	skipNext := false
	result := strings.Builder{}
	utf8RuneArray := []rune(s)
	for i, el := range utf8RuneArray {
		if skipNext {
			skipNext = false
			continue
		}
		if el == '\\' && !asIsNext {
			asIsNext = true
			continue
		}
		if asIsNext && !unicode.IsDigit(el) && el != '\\' {
			return "", ErrInvalidString
		}
		if i+1 < length && unicode.IsDigit(utf8RuneArray[i+1]) {
			if i+2 < length && unicode.IsDigit(utf8RuneArray[i+2]) {
				return "", ErrInvalidString
			}
			counter, _ := strconv.Atoi(string(utf8RuneArray[i+1]))
			result.WriteString(strings.Repeat(string(el), counter))
			skipNext = true
			asIsNext = false
		} else {
			result.WriteRune(el)
			asIsNext = false
		}
	}
	return result.String(), nil
}
