package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/example/stringutil"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(st string) (string, error) {
	if len([]rune(st)) == 0 {
		return st, nil
	}
	st = stringutil.Reverse(st)
	digit_indexes := make(map[int]rune)
	for i, char := range st {
		if unicode.IsDigit(char) {
			digit_indexes[i] = char
		}
	}
	di_slice := make([]int, len(digit_indexes))
	i := 0
	for k := range digit_indexes {
		di_slice[i] = k
		i++
	}
	tmp := -1
	for _, val := range di_slice {
		if val == len(st)-1 {
			return "", ErrInvalidString
		}
		if val-1 == tmp {
			return "", ErrInvalidString
		}
		tmp = val
	}

	var b strings.Builder
	multiplier := 1
	for i, char := range st {
		val, ok := digit_indexes[i]
		if !ok {
			b.WriteString(strings.Repeat(string(char), multiplier))
			multiplier = 1
		}
		if ok {
			multiplier, _ = strconv.Atoi(string(val))
		}
	}
	return stringutil.Reverse(b.String()), nil
}

func main() {
	st := "aaa0b"
	fmt.Println(Unpack(st))
}
