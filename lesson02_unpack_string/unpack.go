package lesson02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	runes := []rune(str)

	var result strings.Builder

	for i := 0; i < len(runes); i++ {

		if i+1 < len(runes) && unicode.IsDigit(runes[i]) && (i == 0 || unicode.IsDigit(runes[i+1])) {

			return "", ErrInvalidString
		} else if i+1 < len(runes) && unicode.IsDigit(runes[i+1]) && unicode.IsLetter(runes[i]) {
			count := int(runes[i+1] - '0')
			letter := string(runes[i])
			newStr := strings.Repeat(letter, count)

			result.WriteString(newStr)
		} else if unicode.IsLetter(runes[i]) {
			letter := string(runes[i])

			result.WriteString(letter)
		}
	}

	return result.String(), nil
}
