package bigint

import (
	"errors"
	"strings"
)

type Sign int

const (
	Positive Sign = iota
	Negative
)

type BigInt struct {
	sign Sign
	num  []int8
}

func (a *BigInt) addToNum(c rune) {
	a.num = append(a.num, int8(c-'0'))
}

func IsInt(s string) bool {
	if len(s) == 0 {
		return false
	}
	if (s[0] != '-' && s[0] != '+') && (s[0] < '0' || s[0] > '9') {
		return false
	}
	for _, c := range s[1:] {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func StrToBigInt(number string) (BigInt, error) {
	number = strings.ReplaceAll(number, " ", "")
	var sign rune
	matched := IsInt(number)
	if !matched {
		return BigInt{}, errors.New("Not a valid number to be converted to BigInt")
	}
	if number[0] == '-' {
		sign = '-'
		number = number[1:]
	} else if number[0] == '+' {
		sign = '+'
		number = number[1:]
	} else {
		sign = '+'
	}
	number = strings.TrimLeft(number, "0")
	if number == "" {
		sign = '+'
		number = "0"
	}
	var num BigInt
	num.num = make([]int8, 0)
	if sign == '-' || sign == '+' {
		if sign == '-' {
			num.sign = Negative
		} else if sign == '+' {
			num.sign = Positive
		}
		for _, c := range number {
			num.addToNum(c)
		}
	} else {
		for _, c := range number {
			num.addToNum(c)
		}
	}
	reverse(num.num)
	return num, nil
}

func (num BigInt) ToString() string {
	var sb strings.Builder
	if num.sign == Negative {
		sb.WriteString("-")
	}
	length := len(num.num)
	for i := length - 1; i >= 0; i-- {
		sb.WriteString(string(num.num[i] + '0'))
	}
	return sb.String()
}
