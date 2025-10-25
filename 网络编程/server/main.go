package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		//创建一个切片，将读取的数据放入切片。
		buf := make([]byte, 1024)
		//从conn读取数据
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("读取失败")
			return
		}
		//将读取的内容输出至终端
		fmt.Println(string(buf[:n]))
	}
}
func main() {
	fmt.Println("服务器端启动")
	lis, err := net.Listen("tcp", "192.168.1.6:8080") //服务器端的ip和端口号
	if err != nil {
		fmt.Println("监听失败", err)
		return
	}
	defer lis.Close()
	//监听成功以后
	//等待客户端的连接
	for {
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println("客户端的等待失败：", err)
		} else {
			fmt.Println("连接成功！")
			fmt.Printf("conn=%v,接收到的客户端信息为：%v\n", conn, conn.RemoteAddr().String())
		}

		//准备协程，写成处理client服务请求
		go process(conn)

	}

}
