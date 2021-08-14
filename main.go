package main

import (
	"github.com/rohith-prakash/mp/test"
)

func main() {
	test.AddTest("1", "1")
	test.AddTest("0", "1")
	test.AddTest("1", "0")
	test.AddTest("-10", "10")
	test.AddTest("1", "-1")
	test.AddTest("-1", "-1")
}
