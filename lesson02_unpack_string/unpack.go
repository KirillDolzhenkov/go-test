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
		if runes[i] == '\\' {
			i++

			if i >= len(runes) {
				return "", ErrInvalidString
			}

			current := runes[i]

			if i+1 < len(runes) && unicode.IsDigit(runes[i+1]) {
				count := int(runes[i+1] - '0')

				result.WriteString(strings.Repeat(string(current), count))
				i++
			} else {
				result.WriteString(string(current))
			}

			continue
		}

		if unicode.IsLetter(runes[i]) {
			current := runes[i]

			if i+1 < len(runes) && unicode.IsDigit(runes[i+1]) {
				count := int(runes[i+1] - '0')

				result.WriteString(strings.Repeat(string(current), count))
				i++
			} else {
				result.WriteString(string(current))
			}
		} else if unicode.IsDigit(runes[i]) {
			if i == 0 {
				return "", ErrInvalidString
			}

			prev := runes[i-1]

			if unicode.IsLetter(prev) == false && unicode.IsDigit(prev) {
				return "", ErrInvalidString
			}
		}
	}

	return result.String(), nil
}

// No asterisk version

//func Unpack(str string) (string, error) {
//	runes := []rune(str)
//
//	var result strings.Builder
//
//	for i := 0; i < len(runes); i++ {
//		if unicode.IsLetter(runes[i]) {
//			current := runes[i]
//
//			if i+1 < len(runes) && unicode.IsDigit(runes[i+1]) {
//				count := int(runes[i+1] - '0')
//
//				result.WriteString(strings.Repeat(string(current), count))
//			} else {
//				result.WriteString(string(current))
//			}
//		} else if unicode.IsDigit(runes[i]) {
//			if i == 0 {
//				return "", ErrInvalidString
//			}
//
//			prev := runes[i-1]
//
//			if unicode.IsLetter(prev) == false && unicode.IsDigit(prev) {
//				return "", ErrInvalidString
//			}
//		}
//	}
//
//	return result.String(), nil
//}
