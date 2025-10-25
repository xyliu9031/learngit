package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup //只定义无需赋值
func main() {

	for i := 0; i < 5; i++ { //启动五个协程
		wg.Add(1) //协程开始的时候加1操作
		go func(name int) {
			defer wg.Done() //防止忘记计数器减一，一般使用defer放在代码首部
			fmt.Println(name)
			// wg.Done() //协程执行完成减1
		}(i)
	}
	wg.Wait() //主协程一直在阻塞，当wg减为0就停止

}
