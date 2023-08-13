package funcs

// 将匿名函数赋值给一个变量
//函数式编程，函数可以作为变量，函数可以作为返回值

var Abc = func() string {
	return "hello"
}

func Fu1() {
	// 匿名函数
	fn := func() string {
		return "hello"
	}
	fn()

	// 匿名函数，立即发起调用,此时fn2不是一个函数，而是该匿名函数的返回值
	fn2 := func() string {
		return "hello"
	}()
	print(fn2)
}

// 返回一个函数

func Fu2() func() string {
	return func() string {
		return "hello"
	}
}

//	函数式编程之闭包
//	闭包=函数+它绑定的上下文
//	闭包如果使用不恰当会导致内存泄露，因为一个对象被闭包引用之后，不能马上被垃圾回收

func Closure(name string) func() string {
	return func() string {
		return "hello" + name
	}
}

// 不定参数，不定参数指函数的最后一个参数可以传入任意多个值，只能将最后一个参数声明为不定参数
// 不定参数在函数内部当作切片来使用    alias := []string{"1","2"}       Yourname("hello",alias...)

func Yourname(name string, alias ...string) {
	if len(alias) > 0 {
		println(alias[0])
	}
}
