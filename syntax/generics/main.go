package main

// T类型参数，名字叫做T，约束是any，等于没有约束

type List[T any] interface {
	add(idx int, t T)
	append(t T)
}

func UseList() {
	var l List[int]
	l.append(1)
}

type LinkedList[T any] struct {
	t T
}
