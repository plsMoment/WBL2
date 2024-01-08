package dev02

import (
	"errors"
	"strconv"
	"unicode"
)

var ErrSyntax = errors.New("first elem must be not a number")

func DecodeRepeats(s string) (string, error) {
	if len(s) < 1 {
		return s, nil
	}
	runeStr := []rune(s)
	if unicode.IsDigit(runeStr[0]) {
		return "", ErrSyntax
	}
	res := make([]rune, 0, len(runeStr))

	left, right := 0, 0
	for left < len(runeStr) {
		if unicode.IsDigit(runeStr[left]) {
			right = left
			for right < len(runeStr) && unicode.IsDigit(runeStr[right]) {
				right++
			}
			num, err := strconv.Atoi(string(runeStr[left:right]))
			if err != nil {
				return "", err
			}

			lastChar := res[len(res)-1]
			for i := 1; i < num; i++ {
				res = append(res, lastChar)
			}
			left = right
		} else {
			res = append(res, runeStr[left])
			left++
		}
	}
	return string(res), nil
}
