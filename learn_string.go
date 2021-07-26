package main

import (
	"fmt"
)

func main()  {
	//str := "go web"
	//fmt.Println(string(str[0])) // 获取索引为0的字符

	//在go语言中，可以用rune和int32来表示一个字符
	//修改字符串的2中方式
	//1、用[]byte修改
	//str := "hi 世界"
	//by := [] byte(str)
	//by[2] = ','
	//fmt.Printf("%s\n", str)
	//fmt.Printf("%s\n", by)

	//2、用[]rune修改
	//str := "hi 世界"
	//by := [] rune(str)
	//by[3] = '中'
	//by[4] = '国'
	//fmt.Println(str)
	//
	//fmt.Println(string(by))

	//单双引号和反引号的问题

	//1、单引号 单引号主要用来为go的特殊类型rune的value赋值
	//rune_test := [] rune("asd")
	//rune_test[1] = 'b'
	//fmt.Println(string(rune_test)) //输出abd
	//
	////2、双引号 双引号通常用来创建字符串变量 支持转义,但是不能用来引用多行
	//a := "aaaa"
	//fmt.Println(a) // 输出aaaa

	//3、反引号 反引号通常用来创建字符串变量 不支持转义，但是支持多行 通常用来书写多行消息、html以及正则等

	//b := `asdasdasdasd
	//		asdasdasdasda` //
	//fmt.Println(b)

	x, y := 6, 8

	fmt.Println(&x)
	fmt.Println(&y)

	change2(&x, &y)

	fmt.Println(x, y)


}
func change2(c, d *int)  {
	fmt.Println(c)
	fmt.Println(d)
	c, d = d, c
	fmt.Println(c)
	fmt.Println(d)
}
