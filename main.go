package main

import (
	"github.com/rohith-prakash/mp/test"
)

func main() {
	test.LogicalOperatorsTest("100", "10")
	test.LogicalOperatorsTest("-100", "100")
	test.LogicalOperatorsTest("-100", "-100")
	test.LogicalOperatorsTest("100", "-100")
	test.LogicalOperatorsTest("-10", "-100")
}
