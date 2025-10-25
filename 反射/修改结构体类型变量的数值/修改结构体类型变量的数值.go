package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func (s Student) Print() {
	fmt.Printf("name:%s age:%d\n", s.Name, s.Age)
}

func (s Student) Set(name string, age int) {
	s.Name = name
	s.Age = age
}

func (s Student) Add(num1 int, num2 int) int {
	return num1 + num2
}

// 定义反射操作的函数
func testReflect(i interface{}) {
	reValue := reflect.ValueOf(i)
	fmt.Printf("reValue:%v\n", reValue)
	n1 := reValue.NumField() //获取有几个字段
	fmt.Printf("n1=%d\n", n1)
	for i := 0; i < n1; i++ {
		fmt.Printf("第%v个字段为：%v\n", i, reValue.Field(i))
	}

	n2 := reValue.NumMethod()
	fmt.Printf("该结构体有%d个方法\n", n2)
	reValue.Method(1).Call(nil) //方法的顺序是按照ASCII码排序

	var nums []reflect.Value
	nums = append(nums, reflect.ValueOf(10))
	nums = append(nums, reflect.ValueOf(20))
	result := reValue.Method(0).Call(nums)
	fmt.Printf("Add加和结果为:%v\n", result[0].Int()) //reValue.Method(i).Call(xxx)返回的是一个切片
}
func main() {
	stu := Student{
		Name: "liuxinyu",
		Age:  25,
	}
	testReflect(stu)

}
