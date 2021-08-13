package test

import (
	"fmt"

	"github.com/rohith-prakash/mp/bigint"
)

func TestHelp(s string) {
	fmt.Println("***TEST START***")
	fmt.Println(s)
	a, err := bigint.StrToBigInt(s)
	if err != nil {
		fmt.Println(err)
		fmt.Println("###TEST END###")
		return
	}
	fmt.Println(a)
	fmt.Println(a.ToString())
	fmt.Println("###TEST END###")
}
