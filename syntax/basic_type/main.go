package main

func main() {
	var a int = 456
	var b int = 123
	println(a + b)
	println(a - b)
	println(a * b)
	if b != 0 {
		println(a / b)
		println(a % b)
	}
	// 同类型才能加减乘除,需要显式类型转换，不支持隐式类型转换
	var c float64 = 12.3
	println(a + int(c))
	var d = 10
	println(d)
	e := 5
	println(e)
	String()
	Byte()
}
