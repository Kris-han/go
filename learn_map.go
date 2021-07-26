package main

//import "fmt"

func main()  {

	//map是一种k v结构的无序集合
	//是引用类型
	//声明方式：var name map[key_type] value_type
	//也可以用 make(map[key_type] value_type)

	//var dict map[int] string
	//dict = map[int]string{0:"kris"}
	//dict := make(map[int] string)
	//dict[0] = "kris"
	//
	//fmt.Println(dict)

	//todo make和new来构造新对象时的区别

	mapTest := new(map[string]string)
	mapTest["k1"] = "k2"

}
