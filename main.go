package main

import (
	"github.com/rohith-prakash/mp/test"
)

func main() {
	test.MagnitudeAddTest("12", "15")
	test.MagnitudeAddTest("121", "15")
	test.MagnitudeAddTest("12", "15")
	test.MagnitudeAddTest("0", "15")
	test.MagnitudeAddTest("15", "15")
	test.MagnitudeAddTest("12", "151")
	test.MagnitudeAddTest("12", "152")
}
