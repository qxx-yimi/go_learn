package main

import "fmt"

// comparable 可比较的    Go在编译的时候，运行的时候，能够判断出来元素是不是相等
// 在switch中，值必须是可比较的
// 在map中，key必须是可比较的
// 基本类型和string也是可比较的
// 如果元素是可比较的，那么该数组也是可比较的

func main() {
	//数组
	a1 := [5]int{1, 2, 3, 4, 5}
	// m的容量为  len(a1)-1 = 5-1 = 4
	m := a1[1:3]
	m[0] = 10 // 会影响a1的值  两者共享内存,   判断子切片会不会对切片/数组产生影响，需要判断两者是否共享内存，如果扩容了，则不共享，不产生影响
	fmt.Printf("m:%v len=%d cap=%d\n", m, len(m), cap(m))
	fmt.Printf("a1:%v len=%d cap=%d\n", a1, len(a1), cap(a1))

	a2 := [4]int{}
	fmt.Printf("a2:%v\n", a2)

	//切片 slice
	a3 := []int{4, 5, 6}
	fmt.Printf("a3:%v len=%d cap=%d\n", a3, len(a3), cap(a3))

	// {0,0,0}  空切片，make([]int,0,4)
	s1 := make([]int, 3, 4)
	s1 = append(s1, 2)
	fmt.Printf("s1:%v len=%d cap=%d\n", s1, len(s1), cap(s1))
	changearr(s1)
	fmt.Printf("s1:%v len=%d cap=%d\n", s1, len(s1), cap(s1))

}

func changearr(a []int) {
	a[0] = 1
	m := make(map[string]int)
	m["1"] = 2
	val, ok := m["hello"]
	if ok {
		print(val) //有对应的键
	}
}
