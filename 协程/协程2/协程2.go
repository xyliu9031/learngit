package main

import (
	"fmt"
	// "strconv"
	"time"
)

func main() { //主线程
	for i := 1; i <= 5; i++ {
		go func(n int) { //匿名函数+外部变量=闭包
			fmt.Println(n)
		}(i)
	}

	time.Sleep(time.Second * 2)

}
