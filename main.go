package main

import (
	"github.com/rohith-prakash/mp/test"
)

func main() {
	test.SubTest("1", "1")
	test.SubTest("0", "1")
	test.SubTest("1", "0")
	test.SubTest("-10", "10")
	test.SubTest("1", "-1")
	test.SubTest("-1", "-1")
}
