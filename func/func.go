package main

func fx(n, m int) int {
	return n + m
}

func fn(n, m int) (r int) {
	r = n + m
	return r
}

func main() {
	println(fx(1, 2))
	println(fx('a', 2))
	println(fn(1, 2))
	// println(fx(false, 2))
}
