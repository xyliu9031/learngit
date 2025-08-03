package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var num int8 = 30
	var num1 float32 = 3.14
	var num2 float32 = 314e-2
	num3 := 3.1415926 //golang中默认为float64类型
	fmt.Println(num, num1, num2, num3)
	fmt.Println(unsafe.Sizeof(num), unsafe.Sizeof(num2))

	var v1 byte = 'a'
	var v2 byte = '6'
	var v3 byte = '('

	fmt.Println(v1, v2, v3)
	fmt.Printf("对应的字符:%c,%c,%c", v1, v2, v3)

	// var s1 string = "xyliu"
	// var s2 = "x"
	// var s3 string = "liuxinyu"
	// var num int8 = 30
	// var num1 float32 = 3.14
	// var num2 float32 = 314e-2
	// num3 := 3.1415926 //golang中默认为float64类型
	// fmt.Println(num, num1, num2, num3)
	// fmt.Println(unsafe.Sizeof(num))

	// var v1 byte = 'a'
	// var v2 byte = '6'
	// var v3 byte = '('

	// fmt.Println(s1, s2)
	// fmt.Println(s3)

	// var s4 string = "liu" + "xin"
	// s5 := s4 + "yu"
	// fmt.Println(s5)

}
