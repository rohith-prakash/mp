package bigint

func reverse(l []uint8) {
	for i, j := 0, len(l)-1; i < j; i, j = i+1, j-1 {
		l[i], l[j] = l[j], l[i]
	}
}

func MagnitudeAdd(a BigInt, b BigInt) BigInt {
	var result []uint8
	var l1, l2 []uint8
	var i = 0
	var carry, sum uint8
	carry, sum = 0, 0
	len_a := len(a.num)
	len_b := len(b.num)
	var len1, len2 int
	if len_a < len_b {
		l1 = make([]uint8, len_a)
		l2 = make([]uint8, len_b)
		copy(l1, a.num)
		copy(l2, b.num)
		len1 = len_a
		len2 = len_b
	} else {
		l1 = make([]uint8, len_b)
		l2 = make([]uint8, len_a)
		copy(l1, b.num)
		copy(l2, a.num)
		len1 = len_b
		len2 = len_a
	}
	reverse(l1)
	reverse(l2)
	//fmt.Print(l1)
	//fmt.Print(l2)
	//fmt.Println(len1 + len2)
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

	reverse(result)
	return BigInt{
		sign: Positive,
		num:  result,
	}
}
