package main

import (
	"fmt"
)

func main() {
	a := Calculator{1}
	// fmt.Println(a.add(2.0))
	// fmt.Println(a.del(1.0))
	// fmt.Println(a.div(0))
	if r, err := a.div(0); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}
	// fmt.Println(mul(&a, 4.0))
}
