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
