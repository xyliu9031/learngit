package main

import (
	"fmt"
	"reflect"
)

//类别 kind  和类型Type不同

func testReflect(i interface{}) { //所有类型都实现了空接口
	//返回Type类型
	reType := reflect.TypeOf(i)
	fmt.Printf("reType:%v\n", reType)
	//fmt.Printf("reType的类型为：%T\n", reType)
	//返回value类型
	reValue := reflect.ValueOf(i)
	fmt.Printf("reValue:%v\n", reValue)
	//fmt.Printf("reValue的类型为：%T\n", reValue)

	//两种方式获取kind类别
	k1 := reType.Kind()
	k2 := reValue.Kind()
	fmt.Println("类别为", k1, k2)

	// reValue转为一个空接口：
	i2 := reValue.Interface()
	n, flag := i2.(Student) //接口使用断言
	if flag {
		fmt.Printf("学生的类型为：%T\n", n)
	}

}

type Student struct {
	Name string
	Age  int
}

func main() {
	stu := Student{"liuxinyu", 25}
	testReflect(stu)

}
