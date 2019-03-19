package main

import "fmt"

func main() {
	a := Calculator{1}
	b := 2.0
	fmt.Println(a.add(b))
	fmt.Println(a.del(1.0))
	fmt.Println(a.div(3.0))
	fmt.Println(mul(&a, 4.0))
}
