package bigint

import (
	"errors"
	"regexp"
	"strings"
)

type Sign int

const (
	Positive Sign = iota
	Negative
)

type BigInt struct {
	sign Sign
	num  []uint8
}

func (a *BigInt) addToNum(c rune) {
	a.num = append(a.num, uint8(c-'0'))
}

func StrToBigInt(number string) (BigInt, error) {
	number = strings.ReplaceAll(number, " ", "")
	matched, err := regexp.MatchString(`[-+]?[0-9]`, number)
	if err != nil {
		return BigInt{}, err
	} else if !matched {
		return BigInt{}, errors.New("Not a valid number to be converted to BigInt")
	}
	var num BigInt
	num.num = make([]uint8, 0)
	if number[0] == '-' || number[0] == '+' {
		if number[0] == '-' {
			num.sign = Negative
		} else if number[0] == '+' {
			num.sign = Positive
		}
		for _, c := range number[1:] {
			num.addToNum(c)
		}
	} else {
		for _, c := range number {
			num.addToNum(c)
		}
	}
	return num, nil
}

func (num BigInt) ToString() string {
	var sb strings.Builder
	if num.sign == Negative {
		sb.WriteString("-")
	}
	for i := range num.num {
		sb.WriteString(string(num.num[i] + '0'))
	}
	return sb.String()
}
