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

type Ordering int

const (
	Lesser Ordering = iota
	Equal
	Greater
	NotEqual
)

type BigInt struct {
	sign Sign
	num  []uint8
}

func (a *BigInt) addToNum(c rune) {
	a.num = append(a.num, uint8(c-'0'))
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
	if number != "0" {
		number = strings.TrimLeft(number, "0")
	} else {
		sign = '+'
	}
	var num BigInt
	num.num = make([]uint8, 0)
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

func MagnitudeCompare(A BigInt, B BigInt) Ordering {
	a := A.num
	b := B.num
	a_len := len(a)
	b_len := len(b)
	if a_len > b_len {
		return Greater
	} else if a_len < b_len {
		return Lesser
	} else {
		for i := 0; i < a_len; i++ {
			if a[i] > b[i] {
				return Greater
			} else if a[i] < b[i] {
				return Lesser
			}
		}
		return Equal
	}
}

func LogicalNegate(order Ordering) Ordering {
	switch order {
	case Lesser:
		return Greater
	case Greater:
		return Lesser
	case Equal:
		return NotEqual
	case NotEqual:
		return Equal
	default:
		return NotEqual
	}
}

func Compare(a BigInt, b BigInt) Ordering {
	if a.sign == Positive && b.sign == Negative {
		return Greater
	} else if a.sign == Negative && b.sign == Positive {
		return Lesser
	} else {
		var magnitudeCompared = MagnitudeCompare(a, b)
		if a.sign == Positive {
			return magnitudeCompared
		} else {
			if magnitudeCompared == Equal {
				return magnitudeCompared
			} else {
				return LogicalNegate(magnitudeCompared)
			}
		}
	}
}
