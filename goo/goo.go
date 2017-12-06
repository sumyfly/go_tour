package main

import (
	"fmt"

	"github.com/dropbox/godropbox/time2"
)

func main() {
	now := time2.NowFloat()
	fmt.Printf("now:%v", now)
}
