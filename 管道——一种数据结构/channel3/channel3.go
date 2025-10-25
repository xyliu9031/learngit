package main

import (
	"fmt"
)

func main() {
	var intChan chan<- int //声明一个管道，具备只写性质
	intChan = make(chan<- int, 3)

	intChan <- 2 //只能写
	// num := <-intChan 不能写

	var intChan2 <-chan int

	if intChan2 != nil {
		fmt.Println(<-intChan2)
	}
}
