package main

import (
	. "fmt"

	qr5 "go_tutor/qr"

	t "github.com/dropbox/godropbox/time2"
)

func main() {
	now := t.NowFloat()
	Printf("now:%v\r\n", now)
	Printf("%v", qr5.QR)
}
