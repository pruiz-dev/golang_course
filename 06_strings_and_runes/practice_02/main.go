package main

import (
	"fmt"
)

func main() {
	str1 := "Пароль"
	s1 := MaskBytesNaive(str1)
	fmt.Printf(
		"s1 is: %v\n\n",
		s1,
	)

	s2, err := MaskRunes(str1, -22, 5, '%')
	if err != nil {
		fmt.Printf("Error: %v\n\n", err)
	} else {
		fmt.Printf(
			"s2 is: %v\n\n",
			s2,
		)
	}

	cardNumber := "4111-2222-3333-4444"
	s3 := MaskCardNumber(cardNumber)
	fmt.Printf(
		"s3 is: %v\n\n",
		s3,
	)

}

func MaskBytesNaive(s string) string {
	bytesArr := []byte(s)
	bytesArr[0] = '*'
	return string(bytesArr)
}

func MaskRunes(s string, start, end int, mask rune) (string, error) {
	var runes []rune = []rune(s)

	if start > end || start < 0 || end > len(runes) {
		return "", fmt.Errorf(
			"ERROR: Incorrect parameters:\n"+
				"start: %d\n"+
				"end: %d\n"+
				"len(runes): %d\n\n",
			start, end, len(runes),
		)
	}

	for i := start; i < end; i++ {
		runes[i] = mask
	}

	return string(runes), nil
}

func MaskCardNumber(card string) string {
	b := []byte(card)
	digitsCount := 0
	for i := range len(b) {
		c := b[i]
		if c >= '0' && c <= '9' {
			digitsCount++
			if digitsCount >= 5 && digitsCount <= 12 {
				b[i] = '*'
			}
		}
	}
	return string(b)
}
