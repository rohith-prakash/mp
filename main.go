package main

import (
	"github.com/rohith-prakash/mp/test"
)

func main() {
	test.TestHelp("0123")
	test.TestHelp("12345")
	test.TestHelp("- 500")
	test.TestHelp("0")
	test.TestHelp("00")
	test.TestHelp("-00")
}
