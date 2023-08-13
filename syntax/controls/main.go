package main

import "fmt"

type User struct {
	name string
}

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	i := 0
	for i < 10 {
		i++
		print(i)
	}

	println("遍历数组")
	arr := [3]string{"11", "22", "33"}
	for i, val := range arr {
		println(i, val)
	}

	println("遍历map")
	m := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	for k, v := range m {
		println(k, v)
	}
	for k := range m {
		println(k, m[k])
	}

	users := []User{
		{
			name: "1",
		},
		{
			name: "2",
		},
	}

	// u 永远是同一个地址,所以不要对迭代参数取地址
	m2 := make(map[string]*User)
	for _, u := range users {
		m2[u.name] = &u
	}

	fmt.Printf("%v", m2)

	//switch 不需要break
	var status = 1
	switch status {
	case 0:
		println("0")
	default:
		println("-1")
	}
}
