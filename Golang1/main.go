// Golang1 project main.go
package main

import (
	f "fmt"
)

func main() {
	//类型断言
	var a interface{} //空接口型变量
	var b int32 = 100 //接口返回的值
	a = b
	f.Println(a)
	c, ok := a.(int32) //需要类型断言，a和c类型不一致 //断言==a.(int)==变量名.(类型)//要赋值的变量什么类型，你就断言什么类型
	//ok是用来判断，断言成功没有？//这是带检测机制的断言
	//同时要使用的变量不能定义数据类型，否则断言检测机制是无用的
	if ok {
		f.Println(c)
	} else {
		f.Println(c, "不是int32型的数值")
	}

}
