package main

import (
	"github.com/rohith-prakash/mp/test"
)

func main() {
	//test.MultiplySimpleTest("933262154439441526816992388562667004907159682643816214685929638952175999932299156089414639761565182862536979208272237582511852109168640000000000000000000000", "100")
	test.AddTest("9", "1")
	test.AddTest("-9", "1")
	test.AddTest("9", "-1")
	test.AddTest("0", "1")
	test.AddTest("1", "-1")
}
