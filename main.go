package main

import (
	"github.com/rohith-prakash/mp/test"
)

func main() {
	test.MagnitudeCompareTest("100", "10")
	test.MagnitudeCompareTest("-100", "100")
	test.MagnitudeCompareTest("-100", "-100")
	test.MagnitudeCompareTest("100", "-100")
	test.MagnitudeCompareTest("-10", "-100")
}
