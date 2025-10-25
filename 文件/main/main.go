package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os" //内置包含文件的标准操作

	"bufio"
)

// IO 即流
func main() {
	file, err := os.Open("E:\\Projects\\go\\hello_world\\main\\test.go")
	if err != nil {
		fmt.Println("打开文件出错，", err)
	}
	fmt.Printf("文件=%v\n", file)
	defer file.Close()
	if err2 := file.Close(); err2 != nil {
		fmt.Println("文件关闭失败！")
	}
	content, err := ioutil.ReadFile("E:\\Projects\\go\\hello_world\\main\\test.go")
	fmt.Printf("%v", string(content))
	fmt.Printf("%v", content)

	//创建一个流，这是当文件较大时采取的操作
	reader := bufio.NewReader(file)
	for {
		std, err := reader.ReadString('\n') //读取到第一换行结束，单个字符用单引号
		if err == io.EOF {                  //表示已经读取到文件结尾
			break
		}
		fmt.Println("读取完成", string(std))
	}

}
