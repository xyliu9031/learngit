package main

import (
	"fmt"
	"strconv"
	"time"
)

// 主死随从
// 主线程中，一个goroutine
func test() {
	for i := 1; i <= 100; i++ {
		fmt.Println("hello golang +" + strconv.Itoa(i))
		time.Sleep(time.Second) //阻塞一秒
	}
}
func main() { //主线程
	go test() //开启一个协程
	for i := 1; i <= 10; i++ {
		fmt.Println("hello msb +" + strconv.Itoa(i))
		time.Sleep(time.Second) //阻塞一秒
	}
}
