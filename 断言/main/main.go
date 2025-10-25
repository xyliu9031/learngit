package main

import (
	"fmt"
)

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

// 实现接口的方法，必须全部实现接口的方法才叫实现
func (p Chinese) sayhello() {
	fmt.Println("你好！")
}
func (p Chinese) chihuoguo() {
	fmt.Println("吃火锅！")
}

type American struct {
	Name string
}

func (p American) sayhello() {
	fmt.Println("hi!")
}

func (p American) chihanbao() {
	fmt.Println("chi han bao!")
}

// 定义一个函数：专门用来各国人打招呼的函数，接收具备SayHello接口的能力的变量：
func greet(s SayHello) { //体现多态
	s.sayhello()

	//断言！！！：
	//var ch Chinese = s.(Chinese) //看s是否能转成chinese类型并且赋值给ch   s.(chinese)表示断言 即类型转换
	// ch, flag := s.(Chinese)  // 返回value,ok  ok是一个bool类型
	// if flag {
	// 	ch.chihuoguo()
	// } else {
	// 	fmt.Println("其他人不会吃火锅")
	// }

	// if ch, flag := s.(Chinese); flag {
	// 	ch.chihuoguo()
	// } else {
	// 	fmt.Println("其他人不会吃火锅")
	// }
	switch s.(type) {
	case Chinese:
		ch := s.(Chinese)
		ch.chihuoguo()
	case American:
		am := s.(American)
		am.chihanbao()

	}

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

	var s SayHello = c
	s.sayhello()

	var s1 SayHello = a
	s1.sayhello()
}
