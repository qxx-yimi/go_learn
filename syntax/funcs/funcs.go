package funcs

// Func1 没有任何参数
func Func1() {

}

func Func2(a int) {

}

func Func3(a int, b string) int {
	return a
}

func Func4(a, b string) {
	// go不支持函数重载
	// 函数式编程，函数可以作为一个变量进行赋值
	myFunc3 := Func3
	_ = myFunc3(1, "2")
}

func Func6(a int, b string) string {
	return "h"
}

func Func7(a int, b string) (string, string) {
	return "a", "b"
}

//函数返回值可以带名字

func Func8(a int, b string) (c int) {
	return 1
}

func func9(a int, b string) (c int) {
	c = 5
	return
}

func func10(a int, b string) (c int) {
	// 将会返回int默认值0
	return
}
