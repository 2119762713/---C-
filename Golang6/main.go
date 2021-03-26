// Golang6 project main.go
package main

import (
	f "fmt"
	"io/ioutil"
	"os"
)

func main() {
	//打开文件
	path := "txt.txt"
	// file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0777) //Opan是只读模式打开
	//openFile是说自己定义打开文件的方式
	//os.O_WRONLY|os.O_APPEND  就是三种方式，是说他是覆盖文件，还是追加，或者只读
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	//关闭文件
	defer file.Close() //延时执行关闭文件
	//写入文件
	//file.WriteString("hello,world")
	//n, err := file.WriteAt([]byte("hello,world."), 0)
	data := bufio.NewReader(path)
	f.Println(n)
	//读取文件
	date, _ := ioutil.ReadFile(path)
	f.Println(date)
}
