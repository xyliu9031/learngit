package main

import (
	"fmt"
	"reflect"
)

func testReflect(i interface{}) { //所有类型都实现了空接口
	//返回Type类型
	reType := reflect.TypeOf(i)
	fmt.Printf("reType:%v\n", reType) //返回value类型
	if "string" == reType.String() {
		fmt.Println("i的类型是string")
	}
	reValue := reflect.ValueOf(i)
	fmt.Printf("reValue:%v\n", reValue)
	fmt.Printf("reValue的类型为：%T\n", reValue)

	//如果想获取reValue的数值
	num2 := 20 + reValue.Int()
	fmt.Println("num2:", num2)

	// reValue转为一个空接口：
	i2 := reValue.Interface()
	n := i2.(int) //接口使用断言
	fmt.Println(n)

}
func main() {
	var num string = "10"
	testReflect(num)

}
