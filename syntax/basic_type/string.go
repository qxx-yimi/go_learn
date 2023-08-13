package main

import "unicode/utf8"

func String() {
	println("Hello \"go")
	println(`
可以换行
不支持转义
再一行
`)

	//字符串拼接,只有string之间才可以直接拼接
	// len() 计算字节长度
	println("hello" + "go")
	println(len("abc"))
	println(len("你好"))
	println(utf8.RuneCountInString("你好abc"))
}
