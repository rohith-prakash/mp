package main

import (
	"github.com/rohith-prakash/mp/test"
)

func main() {
	test.CompareTest("100", "10")
	test.CompareTest("-100", "100")
	test.CompareTest("-100", "-100")
	test.CompareTest("100", "-100")
	test.CompareTest("-10", "-100")
}
