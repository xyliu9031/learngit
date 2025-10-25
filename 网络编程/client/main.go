package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("客户端启动。。。")
	conn, err := net.Dial("tcp", "192.168.1.6:8080") //第一个是传入的协议，第二个是地址端口
	if err != nil {
		fmt.Println("客户端连接失败", err)
		return
	}
	fmt.Println("连接成功，conn:", conn)

	//发送数据，然后退出
	reader := bufio.NewReader(os.Stdin) //os.Stdin待避岙终端标准输入
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("终端输入失败", err)
	}

	//将str发送给服务器
	n, err := conn.Write([]byte(str))
	if err != nil {
		fmt.Println("连接失败", err)
		return
	}
	fmt.Printf("client发送数据成功，一共发送%d字节的数据，并退出", n)
}
