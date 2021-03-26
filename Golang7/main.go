// Golang7 project main.go
package main

import (
	"bufio"
	f "fmt"
	"io/ioutil"
	"os"
)

func main() {
	//使用缓冲区的方式写文件
	//可以有效的保护电脑磁盘
	path := "txt.txt"
	//1.打开文件
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, os.ModePerm) //打开文件
	if err != nil {
		panic(err)
	}

	//2.创建一个内存 ，写指针
	data := bufio.NewWriter(file) //创建文件读指针
	content := "hello,world"
	//延时执行
	defer file.Close() //延时执行关闭文件

	//3.将内容读到内存中
	for _, v := range content {
		_, err1 := data.WriteString(string(v)) //通过WriteString方法以string的类型循环读取
		if err1 != nil {
			panic(err1)
		}
	}
	//清空缓冲
	data.Flush() //写入文件，清空缓冲区
	//读出文件
	res, _ := ioutil.ReadFile(path)
	f.Println("data:", string(res)) //打印
}
