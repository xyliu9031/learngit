package main

import (
	"fmt"
	"sync"
	"time"
)

var num int
var wg sync.WaitGroup
var lock sync.Mutex     //互斥锁,效率相对较低
var rwlock sync.RWMutex //读写锁。读的时候不产生影响，写和读之间才会产生影响

func add() {
	defer wg.Done()
	for i := 1; i <= 100000; i++ {
		lock.Lock()
		num = num + 1
		lock.Unlock()
	}
}
func sub() {
	defer wg.Done()
	for i := 1; i <= 100000; i++ {
		lock.Lock()
		num -= 1
		lock.Unlock()
	}
}

func read() {
	defer wg.Done()
	rwlock.RLock() //如果只读数据，这个锁不产生影响。但是当读写发生时，就会有影响
	fmt.Println("开始读取数据")
	time.Sleep(time.Second)
	fmt.Println("读取成功")
	rwlock.RUnlock()
}

func write() {
	defer wg.Done()
	rwlock.Lock()
	println("开始写入数据")
	time.Sleep(time.Second * 10)
	println("写入完成")
	rwlock.Unlock()
}

func main() {
	wg.Add(2)
	go add() //add和sub里面的同一变量可能争抢资源，因此num最终输出可能不为0
	go sub()
	wg.Wait()
	fmt.Println(num)

	wg.Add(6)
	for i := 1; i <= 5; i++ { //先执行五个读协程
		go read()
	}
	go write() //写少 只有一个协程
	wg.Wait()
	fmt.Println("done")
}

//解决争抢资源的方法:  锁机制——确保一个协程在执行时另外的协程不执行
