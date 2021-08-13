package main

import (
	"fmt"

	"github.com/rohith-prakash/mp/bigint"
	_ "github.com/rohith-prakash/mp/bigint"
)

func main() {
	a := "12345"
	b := "- 12345"
	c, err := bigint.StrToBigInt(a)
	if err != nil {
		fmt.Println(err)
	}
	d, err := bigint.StrToBigInt(b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(c.ToString())
	fmt.Println(d.ToString())
}
