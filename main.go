package main

import (
	"github.com/rohith-prakash/mp/test"
)

func main() {
	test.MultiplySimpleTest("1", "1")
	test.MultiplySimpleTest("0", "1")
	test.MultiplySimpleTest("1", "0")
	test.MultiplySimpleTest("-10", "10")
	test.MultiplySimpleTest("1", "-1")
	test.MultiplySimpleTest("-1", "-1")
	test.MultiplySimpleTest("10", "11")
	test.MultiplySimpleTest("123456", "17")
	test.MultiplySimpleTest("1923456", "1124356")
	test.MultiplySimpleTest("1", "-1")
	test.MultiplySimpleTest("0", "-1")
	test.MultiplySimpleTest("-1", "0")
	test.MultiplySimpleTest("-10", "10")
	test.MultiplySimpleTest("1", "-1")
	test.MultiplySimpleTest("-1", "-1")
	test.MultiplySimpleTest("10", "11")
	test.MultiplySimpleTest("-123456", "17")
	test.MultiplySimpleTest("-1923456", "-1124356")
}
