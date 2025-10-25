package main

import (
	"fmt"
	"sync"
	// "time"
)

var wg sync.WaitGroup

// defer recover
// 输出数字
func printNum() {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func devide() {
	defer wg.Done()
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("发生错误", err)
		}
	}()
	num1 := 10
	num2 := 0

	fmt.Println(num1 / num2) //0作为除数会报错
}
func main() {
	wg.Add(2)
	go printNum()
	go devide()
	wg.Wait()

}
