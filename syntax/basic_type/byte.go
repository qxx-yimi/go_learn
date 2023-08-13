package main

import "fmt"

func Byte() {
	// byte 是uint8的别名
	var a byte = 'a'
	//打印ascii码 97
	println(a)
	fmt.Printf("%c\n", a)

	//byte 和 string
	var str string = "this is string"
	var bs []byte = []byte(str)
	println(str, bs)
}
