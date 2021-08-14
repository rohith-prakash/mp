package bigint

import (
	"errors"
)

func reverse(l []int8) {
	for i, j := 0, len(l)-1; i < j; i, j = i+1, j-1 {
		l[i], l[j] = l[j], l[i]
	}
}

func MagnitudeAdd(a BigInt, b BigInt) BigInt {
	var result []int8 = make([]int8, 0)
	var l1, l2 []int8
	var i = 0
	var carry, sum int8
	carry, sum = 0, 0
	len_a := len(a.num)
	len_b := len(b.num)
	var len1, len2 int
	if len_a < len_b {
		l1 = a.num
		l2 = b.num
		len1 = len_a
		len2 = len_b
	} else {
		l1 = b.num
		l2 = a.num
		len1 = len_b
		len2 = len_a
	}

	for i = 0; i < len1; i++ {
		sum = l1[i] + l2[i] + carry
		result = append(result, sum%10)
		carry = sum / 10
	}
	for i = len1; i < len2; i++ {
		sum = l2[i] + carry
		result = append(result, sum%10)
		carry = sum / 10
	}

	return BigInt{
		sign: Positive,
		num:  result,
	}
}

//Unless |a|>= |b| don't use this
func MagnitudeSub(a BigInt, b BigInt) (BigInt, error) {
	var result []int8 = make([]int8, 0)
	var carry, diff int8
	var i int
	l1, l2 := a.num, b.num
	len1 := len(l1)
	len2 := len(l2)
	// if len2 > len1 {
	// 	return BigInt{}, errors.New("Numer with larger magnitude must be first argument")
	// }
	if MagnitudeCompare(a, b) == Lesser {
		return BigInt{}, errors.New("Numer with larger magnitude must be first argument")
	}
	carry = 0
	for i = 0; i < len2; i++ {
		diff = ((l1[i] - l2[i]) - carry)
		if diff < 0 {
			diff += 10
			carry++
		} else {
			carry = 0
		}
		result = append(result, diff)
	}
	for i = len2; i < len1; i++ {
		diff = l1[i] - carry
		if diff < 0 {
			diff += 10
			carry++
		} else {
			carry = 0
		}
		result = append(result, diff)
	}
	answer := BigInt{
		sign: Positive,
		num:  result,
	}
	answer.Clean()
	return answer, nil
}

func (a BigInt) checkIfAllZero() bool {
	for i := range a.num {
		if a.num[i] != 0 {
			return false
		}
	}
	return true
}

func (a *BigInt) Clean() {
	if a.checkIfAllZero() {
		a.sign = Positive
		a.num = []int8{0}
		return
	}
	var numOfZeros int = 0
	length := len(a.num)
	for i := length - 1; i > 0; i-- {
		if a.num[i] != 0 {
			break
		}
		numOfZeros++
	}
	a.num = a.num[:length-numOfZeros]
}

// func Add(a BigInt,b BigInt) BigInt {
// 	c := MagnitudeAdd(a,b)

// }
