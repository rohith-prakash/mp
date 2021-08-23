package test

import (
	"fmt"

	"github.com/rohith-prakash/mp/bigint"
)

func TestHelp(s string) {
	fmt.Println("***TEST START***")
	fmt.Println(s)
	a, err := bigint.StrToBigInt(s)
	if err != nil {
		fmt.Println(err)
		fmt.Println("###TEST END###")
		return
	}
	fmt.Println(a)
	fmt.Println(a.ToString())
	fmt.Println("###TEST END###")
}

func MagnitudeCompareTest(a string, b string) {
	a2, err := bigint.StrToBigInt(a)
	if err != nil {
		fmt.Println("A cannot be converted")
		return
	}
	b2, err := bigint.StrToBigInt(b)
	if err != nil {
		fmt.Println("A cannot be converted")
		return
	}
	switch bigint.MagnitudeCompare(a2, b2) {
	case bigint.Lesser:
		fmt.Println(a, "<", b)
	case bigint.Equal:
		fmt.Println(a, "==", b)
	case bigint.Greater:
		fmt.Println(a, ">", b)
	default:
		fmt.Println("Unexpected outcome")
	}
}

func CompareTest(a string, b string) {
	a2, err := bigint.StrToBigInt(a)
	if err != nil {
		fmt.Println("A cannot be converted")
		return
	}
	b2, err := bigint.StrToBigInt(b)
	if err != nil {
		fmt.Println("A cannot be converted")
		return
	}
	switch bigint.Compare(a2, b2) {
	case bigint.Lesser:
		fmt.Println(a, "<", b)
	case bigint.Equal:
		fmt.Println(a, "==", b)
	case bigint.Greater:
		fmt.Println(a, ">", b)
	default:
		fmt.Println("Unexpected outcome")
	}
}

func LogicalOperatorsTest(a string, b string) {
	a2, err := bigint.StrToBigInt(a)
	if err != nil {
		fmt.Println("A cannot be converted")
		return
	}
	b2, err := bigint.StrToBigInt(b)
	if err != nil {
		fmt.Println("A cannot be converted")
		return
	}
	if bigint.LessThan(a2, b2) {
		fmt.Println(a, " is less than ", b)
	} else if bigint.GreaterThan(a2, b2) {
		fmt.Println(a, " is greater than ", b)
	} else if bigint.EqualTo(a2, b2) {
		fmt.Println(a, " is equal to ", b)
	}
}

func MagnitudeAddTest(a string, b string) {
	a2, err := bigint.StrToBigInt(a)
	if err != nil {
		fmt.Println("A cannot be converted")
		return
	}
	b2, err := bigint.StrToBigInt(b)
	if err != nil {
		fmt.Println("A cannot be converted")
		return
	}
	c := bigint.MagnitudeAdd(a2, b2)
	fmt.Println(a2.ToString(), "+", b2.ToString(), "=", c.ToString())
}

func MagnitudeSubTest(a string, b string) {
	a2, err := bigint.StrToBigInt(a)
	if err != nil {
		fmt.Println("A cannot be converted")
		return
	}
	b2, err := bigint.StrToBigInt(b)
	if err != nil {
		fmt.Println("A cannot be converted")
		return
	}
	c, err := bigint.MagnitudeSub(a2, b2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(a2.ToString(), "-", b2.ToString(), "=", c.ToString())
}

func CleanTest(a string, b string) {
	a2, err := bigint.StrToBigInt(a)
	if err != nil {
		fmt.Println("A cannot be converted")
		return
	}
	b2, err := bigint.StrToBigInt(b)
	if err != nil {
		fmt.Println("A cannot be converted")
		return
	}
	c, err := bigint.MagnitudeSub(a2, b2)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.Clean()
	fmt.Println(a2.ToString(), "-", b2.ToString(), "=", c.ToString())
}

func AddTest(a string, b string) {
	a2, err := bigint.StrToBigInt(a)
	if err != nil {
		fmt.Println("A cannot be converted")
		return
	}
	b2, err := bigint.StrToBigInt(b)
	if err != nil {
		fmt.Println("A cannot be converted")
		return
	}
	c := bigint.Add(a2, b2)
	fmt.Println(a2.ToString(), "+", b2.ToString(), "=", c.ToString())
}

func SubTest(a string, b string) {
	a2, err := bigint.StrToBigInt(a)
	if err != nil {
		fmt.Println("A cannot be converted")
		return
	}
	b2, err := bigint.StrToBigInt(b)
	if err != nil {
		fmt.Println("A cannot be converted")
		return
	}
	c := bigint.Sub(a2, b2)
	fmt.Println(a2.ToString(), "-", b2.ToString(), "=", c.ToString())
}

func MagnitudeMultiplySimpleTest(a string, b string) {
	a2, err := bigint.StrToBigInt(a)
	if err != nil {
		fmt.Println("A cannot be converted")
		return
	}
	b2, err := bigint.StrToBigInt(b)
	if err != nil {
		fmt.Println("A cannot be converted")
		return
	}
	c := bigint.MagnitudeMultiplySimple(a2, b2)
	fmt.Println(a2.ToString(), "*", b2.ToString(), "=", c.ToString())
}
