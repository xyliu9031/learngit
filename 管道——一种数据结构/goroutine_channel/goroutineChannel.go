package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func writeData(intChan chan int) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		intChan <- i
		fmt.Println("写入数据——", i)
		time.Sleep(time.Second)
	}
	close(intChan)
}
func readeData(intChan chan int) {
	defer wg.Done()
	for i := range intChan {
		fmt.Println("读取数据", i)
		time.Sleep(time.Second)
	}
}
func main() {
	//写goroutine和读goroutine共同操作一个管道
	intChan := make(chan int, 5)
	wg.Add(2)
	go writeData(intChan)
	go readeData(intChan)
	wg.Wait()

}
