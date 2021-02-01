// Golang2 project main.go
package main

import (
	"database/sql"
	f "fmt"
	"net/http"
	"html/template"
	_ "github.com/go-sql-driver/mysql"
)

func baidu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "text/html") //响应头
	w.Write([]byte('<html><head><title>分类列表</title></head>'))
	w.Write([]byte("<body>"))
	Db,err := sql.Open('mysql','root:root@tcp(127.0.0.1:3306)/bak?charset=utf8;')
	res,err:=DB.Query('select * from yunwu')
	if != nil{
		w.write([]byte("没有数据"))
	}
	return
}

func index(w http.ResponseWriter, r *http.Request){
	file,err := template.ParseFiles("./html/index")
	if != nil{
		f.Printf(err)
		w.write([]byte("404 not found"))
		return
	}
	//str := "1904班"
	type str 
	file.Execute(w,str)
}
func main() {
	http.HandleFunc("/hello", lit)
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}
