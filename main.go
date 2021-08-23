package main

import (
	"github.com/rohith-prakash/mp/test"
)

func main() {
	test.MagnitudeMultiplySimpleTest("1", "1")
	test.MagnitudeMultiplySimpleTest("0", "1")
	test.MagnitudeMultiplySimpleTest("1", "0")
	test.MagnitudeMultiplySimpleTest("-10", "10")
	test.MagnitudeMultiplySimpleTest("1", "-1")
	test.MagnitudeMultiplySimpleTest("-1", "-1")
	test.MagnitudeMultiplySimpleTest("10", "11")
	test.MagnitudeMultiplySimpleTest("123456", "17")
	test.MagnitudeMultiplySimpleTest("1923456", "1124356")
}
