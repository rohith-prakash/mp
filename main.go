package main

import (
	"github.com/rohith-prakash/mp/test"
)

func main() {
	test.TestHelp("12345")
	test.TestHelp("+12345")
	test.TestHelp("-12345")
	test.TestHelp("- 12345")
	test.TestHelp("1+2345")
	test.TestHelp("*12345")
}
