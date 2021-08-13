package bigint

type Ordering int

const (
	Lesser Ordering = iota
	Equal
	Greater
	NotEqual
)

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
		for i := a_len - 1; i >= 0; i-- {
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

func LessThan(a BigInt, b BigInt) bool {
	return Compare(a, b) == Lesser
}

func GreaterThan(a BigInt, b BigInt) bool {
	return Compare(a, b) == Greater
}

func EqualTo(a BigInt, b BigInt) bool {
	return Compare(a, b) == Equal
}

func NotEqualTo(a BigInt, b BigInt) bool {
	return Compare(a, b) != Equal
}
