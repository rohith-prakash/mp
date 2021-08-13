package main

import (
	"github.com/rohith-prakash/mp/test"
)

func main() {
	test.MagnitudeSubTest("15", "12")
	test.MagnitudeSubTest("121", "15")
	test.MagnitudeSubTest("120", "15")
	test.MagnitudeSubTest("0", "0")
	test.MagnitudeSubTest("15", "15")
	test.MagnitudeSubTest("1222", "151")
	test.MagnitudeSubTest("1521", "152")
}
