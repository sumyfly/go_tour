package main

import "fmt"

const (
	A = iota + 1
	B
)

func AX() int {
	return A + B
}

func main() {
	fmt.Println(A)
	fmt.Println(B)
}
