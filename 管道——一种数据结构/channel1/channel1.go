package main

import (
	"fmt"
)

func main() {
	var intChan chan int
	intChan = make(chan int, 3) //管道可以存放3个int类型的数据

	intChan <- 10
	num := 20
	intChan <- num
	fmt.Printf("管道的实际长度：%v,管道的容量是：%v", len(intChan), cap(intChan))

	num1 := <-intChan
	fmt.Println("\n", num1)
	close(intChan)
	// intChan <- 30
	num2 := <-intChan
	fmt.Println(num2)

	var intChan2 chan int
	intChan2 = make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i
	}
	// 在遍历前，如果没有关闭管道，就会报deadlock错误
	close(intChan2)
	for value := range intChan2 {
		fmt.Println("value =", value)
	}

}
