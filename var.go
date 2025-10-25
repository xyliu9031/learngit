package main

import "fmt"

//全局变量的声明
var (
	n9  int    = 500
	n10 string = "asdaq"
)

func main() {

	//局部变量
	var num int = 18
	fmt.Println(num)

	var num2 int
	fmt.Println(num2)

	var num3 = 12.5

	sex := "male" //省略var
	var n1, n2, n3 int
	fmt.Println(n1, n2, n3)
	var n4, name, n5 = 10, "xyliu", 7.8

	n7, n8 := 123, "asdas"

}

/*基本数据类型*/
/*
数值型
字符型 一个字节 byte  直接输出为对应的ASCII ，必须用%c才可以输出原来的字符
布尔型 bool 占一个字节 true 和false
字符串
*/

/*派生数据类型/复杂数据类型*/
/*
指针、数组、结构体、管道、函数、切片、接口、map
*/
