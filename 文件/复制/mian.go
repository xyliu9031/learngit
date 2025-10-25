package main

import (
	"fmt"
	"os"
)

func main() {
	file1 := "E:\\Projects\\go\\hello_world\\main\\test.go"
	file2 := "E:\\Projects\\go\\hello_world\\main\\test1.go"

	content, err := os.ReadFile(file1)
	if err != nil {
		fmt.Println("读取失败")
		return
	}

	err = os.WriteFile(file2, content, 0666)
	if err != nil {
		fmt.Println("写出失败")

	}

}
