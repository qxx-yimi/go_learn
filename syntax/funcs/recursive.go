package funcs

//递归一定要有中断条件，一定会结束，不然会导致栈溢出

func Recursive() {
	Recursive()
}

func A() {
	B()
}

func B() {
	C()
}

func C() {
	A()
}
