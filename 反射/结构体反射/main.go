package main

import (
	"fmt"
	"reflect"
)

func testReflect(i interface{}) { //所有类型都实现了空接口
	//返回Type类型
	reType := reflect.TypeOf(i)
	fmt.Printf("reType:%v\n", reType)
	fmt.Printf("reType的类型为：%T\n", reType)
	//返回value类型
	reValue := reflect.ValueOf(i)
	fmt.Printf("reValue:%v\n", reValue)
	fmt.Printf("reValue的类型为：%T\n", reValue)

	// reValue转为一个空接口：
	i2 := reValue.Interface()
	n, flag := i2.(Student) //接口使用断言
	if flag {
		fmt.Printf("学生的名字为：%v。学生的年龄为：%v\n", n.Name, n.Age)
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
