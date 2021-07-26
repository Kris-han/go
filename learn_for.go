package main

import "fmt"



func main()  {

	//go中的for写法类似js
	//for 初始化语句; 条件语句; 修饰语句 {}

	//for i :=0; i< 10;i ++{
	//	fmt.Printf("i is %d", i)
	//}

	//===================================
	//===================================

	//go中的第二种写法，类似pyton的 while写法
	//for 条件语句 {}
	//var i int = 10
	//for i >= 0{
	//	i -= 1
	//	fmt.Println("i is, %d", i)
	//}


	//=====================================
	//=====================================

	//go中的无限循环 死循环,类似python的 while true
	//for { }
	//var a int = 1
	//for{
	//	if a >= 5{
	//		break
	//	}
	//	a += 1
	//	fmt.Println("a is, %d", a)
	//}


	//=====================================
	//=====================================

	//go中的for range  这是 Go 特有的一种的迭代结构
	//它可以迭代任何一个集合（包括数组和 map，详见第 7 和 8 章）。语法上很类似其它语言中 foreach 语句，
	//但您依旧可以获得每次迭代所对应的索引。一般形式为：for ix, val := range coll { }
	//类似python中的 for key, value in dict().item():

	//str := "Go is a beautiful language!"
	//fmt.Println("str length is, %d", len(str))
	//var  x  string = "2"
	//for key, value := range str{
	//	fmt.Printf("key is %d, value is %c ||||", key, value)
	//	fmt.Println("key is %d, value is %c", key, value)
	//
	//	fmt.Println("x is %c:", x)
	//}

	//另外：Printf是格式化输出，以后格式化输出用这个 value 是变量，
	//剩下的用 Println value 是 改变量的unicode码

	//break 退出当前循环 和python用法一样

	//continue 忽略当前循环, 直接执行下一次循环 和python用法一样 只能用于for 循环中

	// 标签:  for、switch 或 select 语句都可以配合标签（label）形式的标识符使用，
	//		 即某一行第一个以冒号（:）结尾的单词（gofmt 会将后续代码自动移至下一行）。
LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}

}


