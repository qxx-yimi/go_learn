package main

import (
	"fmt"
)

// 接口是一组行为的抽象，面向接口编程
// 接口名的首字母和其中元素的首字母大小写都会影响可访问性

type List interface {
	Add(index int, val any) error
	Append(val any)
	Delete(index int) (any, error)
}

// 结构体初始化，Go没有构造函数，初始化语法 LinkList{}, 获取指针&LinkList{}，new(LinkList)
// 当一个结构体具备接口的全部方法时，它就实现了该接口
// 实现了接口的结构体，可以赋值给该接口

type LinkList struct {
	head *node
	tail *node
	Len  int
}

func (l LinkList) Append(val any) {
	//TODO implement me
	panic("implement me")
}

func (l LinkList) Delete(index int) (any, error) {
	//TODO implement me
	panic("implement me")
}

type node struct {
}

//方法接收器receiver

func (l LinkList) Add(index int, val any) {

}

// 声明成指针类型，才能改变

func (l *LinkList) AddV2(index int, val any) {

}

type User struct {
	Name string
}

func main() {
	u1 := &User{}
	println(u1)

	u2 := User{Name: "hello"}
	println(u2.Name)

	u3 := new(User)
	println(u3)

	fmt.Printf("%+v\n", u2)
	fmt.Printf("%v\n", u2)
}

// 衍生类型，type TypeA TypeB
// 衍生类型是一个全新的类型，TypeB实现了某个接口，不等于TypeA也实现了某个接口，衍生类型之间可以通过（）进行转换
// type Integer int
// i1 := 10
// i2 := Integer(i1)
// 别名 type TypeA = TypeB   一模一样，换个名字而已

type Fish struct {
	Name string
}

func (f Fish) Swim() {
	println("fish is swimming")
}

type FakeFish Fish

func UseFish() {
	//f1 := Fish{}
	//f2 := FakeFish{(f1)}
	// f2.Swim()   f2 无Swim方法   但是f2.Name可以正常使用
}

// 组合
// 结构体组合结构体，接口组合接口，结构体组合结构体指针，
