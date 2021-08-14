package main

import (
	"github.com/rohith-prakash/mp/test"
)

func main() {
	test.MagnitudeSubTest("-15", "12")
	test.MagnitudeSubTest("121", "-121")
	test.MagnitudeSubTest("0", "0")
	test.MagnitudeSubTest("010", "10")
	test.MagnitudeSubTest("105", "1050")
	test.MagnitudeSubTest("1222", "151")
	test.MagnitudeSubTest("1521", "152")
}
