package main

import (
	"errors"
	"fmt"
)

func Delete[T any](slice []T, idx int) ([]T, error) {
	if idx < 0 || idx > len(slice)-1 {
		return nil, errors.New("index is wrong")
	}
	for i := idx; i+1 < len(slice); i++ {
		slice[i] = slice[i+1]
	}
	return slice[:len(slice)-1], nil
}

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice, _ = Delete[int](slice, 7)
	fmt.Printf("%v", slice)
}
