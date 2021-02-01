// Golang1 project main.go
package main

import (
	//f "fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func Hello1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("第一个页面"))
}

func Hello2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("第二个页面"))
}

func Hello3(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("第三个页面"))
}

func main() {
	http.HandleFunc("/1", Hello1)
	http.HandleFunc("/2", Hello2)
	http.HandleFunc("/3", Hello3)
	http.ListenAndServe(":8000", nil)
}
