package main

// defer 配合  闭包  使用
// 输出 1

func DeferClosure() {
	i := 0
	defer func() {
		println(i)
	}()
	i = 1
}

// 输出 0  defer传参

func DeferClosureV2() {
	i := 0
	defer func(val int) {
		println(val)
	}(i)
	i = 1
}

func DeferClosureLoopV1() {
	for i := 0; i < 10; i++ {
		defer func() {
			println(i)
		}()
	}
}

func DeferClosureLoopV2() {
	for i := 0; i < 10; i++ {
		defer func(val int) {
			println(val)
		}(i)
	}
}

func DeferClosureLoopV3() {
	for i := 0; i < 10; i++ {
		j := i
		defer func() {
			println(j)
		}()
	}
}

func main() {
	DeferClosure()       // 1
	DeferClosureV2()     // 0
	DeferClosureLoopV1() // 10 10 10 10 10 10 10 10 10 10
	DeferClosureLoopV2() // 9  8  7  6  5  4  3  2  1  0
	DeferClosureLoopV3() // 9  8  7  6  5  4  3  2  1  0
}
