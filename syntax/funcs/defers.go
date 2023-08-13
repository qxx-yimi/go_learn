package funcs

//defer 在函数返回的前一刻，执行一段逻辑，defer：延迟调用
// 该函数会输出 21 而不是 12，类似于栈  即defer的执行顺序，在逻辑上与栈相同

func Defer() {
	defer func() {
		println("1")
	}()

	defer func() {
		println("2")
	}()
}
