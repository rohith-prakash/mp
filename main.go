package main

import (
	"fmt"

	"github.com/rohith-prakash/mp/test"
)

func main() {
	test.MagnitudeCompareTest("1200", "1021")
	test.MagnitudeCompareTest("- 1200", "- 1021")
	test.MagnitudeCompareTest("-1000", "0")
	test.MagnitudeCompareTest("-0", "0")
	test.MagnitudeCompareTest("-12", "12")
	test.MagnitudeCompareTest("0", "1021")
	test.MagnitudeCompareTest("-1200", "1021")
	fmt.Println("##################")
	test.CompareTest("1200", "1021")
	test.CompareTest("- 1200", "- 1021")
	test.CompareTest("-1000", "0")
	test.CompareTest("-0", "0")
	test.CompareTest("-12", "12")
	test.CompareTest("0", "1021")
	test.CompareTest("-1200", "1021")
}
