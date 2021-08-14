package main

import (
	"github.com/rohith-prakash/mp/test"
)

func main() {
	test.CleanTest("15", "12")
	test.CleanTest("121", "121")
	test.CleanTest("0", "-0")
	test.CleanTest("010", "10")
	test.CleanTest("105", "105")
	test.CleanTest("1222", "151")
	test.CleanTest("1521", "152")
}
