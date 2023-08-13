package main

func Sum[T Number](vals ...T) T {
	var res T
	for _, val := range vals {
		res += val
	}
	return res
}

// Number是泛型约束

type Number interface {
	int | int64 | float64
}

func main() {
	println(Sum[int](1, 2, 3))
}
