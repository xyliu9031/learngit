package main

import (
	"fmt"
)

// 一个变量可以实现多个接口
// 接口的定义：定义规则、定义规范
// 接口定义一组方法 但不需要实现，实现定义结构体来做，或者自定义类型  例如type inter int   ，inter类型也可以实现
// 不实现接口，则使用时返回一个nil
type SayHello interface {
	//声明没有实现的方法：
	sayhello()
}

// 接口的实现：定义结构体
type Chinese struct {
	Name string
}

// 定义一个函数：专门用来各国人打招呼的函数，接收具备SayHello接口的能力的变量：
func greet(s SayHello) { //体现多态，多态是通过接口实现的。Chinese和American通过SayHello这个接口实现
	s.sayhello()
}

// 实现接口的方法，必须全部实现接口的方法才叫实现
func (p Chinese) sayhello() {
	fmt.Println("你好！")
}

type American struct {
	Name string
}

func (p American) sayhello() {
	fmt.Println("hi!")
}

type kongjiekou interface {
}

func main() {
	// 定义一个SayHello接口数组，里面存放American，Chinese结构体变量
	var arr [3]SayHello //多态数组
	arr[0] = American{Name: "jack"}
	arr[1] = Chinese{Name: "刘欣雨"}
	fmt.Println(arr)

	c := Chinese{}
	a := American{}
	greet(c)
	greet(a)
	c.sayhello()
	a.sayhello()

	var s SayHello = c //接口本身不能创建实例，但是可以指向一个实现接口功能的变量
	s.sayhello()
	var s1 SayHello = a
	s1.sayhello()

	// 任何其他类型可以实现空接口
	var ch Chinese
	var e kongjiekou = ch
	fmt.Println(e)

	var num float64 = 9.3
	var e1 interface{} = num
	fmt.Println(e1)
}
