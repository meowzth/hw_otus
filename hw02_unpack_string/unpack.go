package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/example/stringutil"
)

var ErrInvalidString = errors.New("invalid string")

var ErrAmbiquosString = errors.New("ambiquos string")

func Unpack(st string) (string, error) {
	if len([]rune(st)) == 0 {
		return st, nil
	}
	st = stringutil.Reverse(st)

	digitIndx := make(map[int]rune)
	for i, char := range st {
		if unicode.IsDigit(char) {
			digitIndx[i] = char
		}
	}
	digitIndxSlice := make([]int, len(digitIndx))
	i := 0
	for k := range digitIndx {
		digitIndxSlice[i] = k
		i++
	}
	tmp := -2
	for _, val := range digitIndxSlice {
		if val == len(st)-1 {
			if val == 0 {
				return st, ErrAmbiquosString
			}
			return "", ErrInvalidString
		}
		if val-1 == tmp {
			return "", ErrInvalidString
		}
		tmp = val
	}

	var res strings.Builder
	multiplier := 1
	for i, char := range st {
		val, ok := digitIndx[i]
		if !ok {
			res.WriteString(strings.Repeat(string(char), multiplier))
			multiplier = 1
		}
		if ok {
			multiplier, _ = strconv.Atoi(string(val))
		}
	}
	return stringutil.Reverse(res.String()), nil
}
