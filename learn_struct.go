package main

import "fmt"

func main(){


	//数组的生命方式：var name[size] type   通过下标索引取数
	//var name_list[3] string

	//初始化方式：在经过初始化的数组中，{}中的元素个数不能大于[]的数字，默认情况下，如果无法确定数组大小，可以用[...]来确定
	//var init_list = [3]string{"a", "b", "c"}
	//var init_list2 = [...]string{"a", "b", "c", "x"}


	//结构体 结构体必须有一下特征：
	//1、字段拥有自己的类型和值
	//2、字段名必须唯一
	//3、字段的类型也可以是结构体，甚至是字段所在结构体的类型
	//定义格式如下：
	/*type 类型名 struct {
		字段1 类型1
		字段2 类型2
	}*/
	//type Pointer struct {
	//	A float32
	//	B float64
	//}
	//
	//type color struct {
	//	red, blue, green byte
	//}




	//book := Book{"golang学习笔记", 10, "kris"}
	//printBook(book)

	//结构体指针：
	//声明方式:
	//var name *指针类型结构体
	var structPointer Book
	structPointer.name = "golang学习笔记"
	structPointer.price = 12
	structPointer.author = "kris2"
	printBook(&structPointer)



}
//demo
func printBook(book *Book){
	fmt.Println(book.name)
	fmt.Println(book.price)
	fmt.Println(book.author)
}
type Book struct {
	name string
	price int
	author string
}
