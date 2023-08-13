package main

var Global = "全局变量，包外也可见，首字母必须大写"
var internal = "包内变量，私有变量，包内可见，首字母不能大写"

func main() {
	// 局部变量
	//int 是灰色的，因为数字默认是int类型,go支持类型推断
	var a int = 123
	println(a)

	var b = 234
	println(b)

	var (
		d string = "123"
		e int    = 5
	)
	println(d)
	println(e)

	//该方式只能用于局部变量
	f := 123
	println(f)

	// := 左侧至少要有一个新变量
}
