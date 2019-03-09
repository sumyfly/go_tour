package main

import "fmt"

func sum(arrays []int, ch chan int) {
	//fmt.Println(arrays)
	sum := 0
	for _, array := range arrays {
		sum += array
	}
	ch <- sum
}

func recv(arrayChan chan int, arrayResult [10]int) {
	for i := 0; i < 10; i++ {
		arrayResult[i] = <-arrayChan
	}
	fmt.Println(arrayResult)
}

func main() {
	arrayChan := make(chan int, 0)
	arrayResult := [10]int{0}

	arrayInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	length := len(arrayInt)
	for t := 0; t < 10; t++ {
		go sum(arrayInt[length-t:], arrayChan)
	}
	recv(arrayChan, arrayResult)
}
