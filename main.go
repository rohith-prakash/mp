package main

import (
	"github.com/rohith-prakash/mp/test"
)

func main() {
	test.MagnitudeAddTest("123", "456")
	test.MagnitudeAddTest("0", "123")
	test.MagnitudeAddTest("9", "21")
}
