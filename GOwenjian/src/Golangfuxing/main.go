// Golangfuxing project main.go
package main

import (
	f "fmt"
)

func main() {
	var a [3]uint
	a[0] = 100
	a[1] = 200
	a[2] = 300
	for i := 0; i < len(a); i++ {
		f.Println(a[i])
	}
	f.Println("//==========================================================================================")
	for _, v := range a {
		f.Println(v)
	}
	f.Println("//==========================================================================================")
	type p struct {
		p1 string
	}
	type aaa struct {
		a1 int
		a2 byte
		p
	}
	f.Println("//==========================================================================================")
}
