package hw02unpackstring

import (
	"errors"
	"strconv"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	runes := []rune(s)
	var outString string
	for i := 0; i < len(runes); i++ {
		currentChar := runes[i]

		switch {
		case currentChar > 47 && currentChar < 58 && i == 0: // первая цифра в строке
			return "", ErrInvalidString
		case currentChar > 47 && currentChar < 58 && runes[i-1] > 47 && runes[i-1] < 58: // предыдущий символ тоже цифра
			return "", ErrInvalidString
		case currentChar == 48: // zero
			outString = string([]rune(outString)[:len([]rune(outString))-1])
			continue
		case currentChar > 48 && currentChar < 58: // цифра от 1 до 9
			digit, _ := strconv.Atoi(string(currentChar))
			for x := 0; x < digit-1; x++ {
				outString += string(runes[i-1])
			}
			continue
		}

		outString += string(currentChar)
	}
	return outString, nil
}
