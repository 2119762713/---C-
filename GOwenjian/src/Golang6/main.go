// Golang6 project main.go
package main

import (
	_ "database/sql"
	f "fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func index(w http.ResponseWriter, r *http.Request) {
	file, err := template.ParseFiles("./html/index.html")
	if err != nil {
		f.Println(err) //不会显示在网页
		w.Write([]byte("404 not found"))
		return
	}
	file.Execute(w, nil)
	//==============================
	// Db, err2 := sql.Open("mysql", `root:root@tcp(127.0.0.1:3306)/bak?charset=utf8;`)
	// if err2 != nil {
	// 	f.Println("数据连接失败", err2)
	// }
	// res, _ := Db.Query("select id,name from yunwu")

	// var a []map[string]string
	// for res.Next() {
	// 	var data map[string]string = make(map[string]string)
	// 	var id string
	// 	var name string
	// 	res.Scan(&id, &name)
	// 	data["id"] = id
	// 	data["name"] = name
	// 	a = append(a, data)
	// }
	// type Vars struct {
	// 	Date []map[string]string
	// }
	// var aaa Vars = Vars{a}
	// file.Execute(w, aaa)
	//==============================
	// var slice []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// var b bool = true
	// type Vars struct {
	// 	Name  string
	// 	Slice []int
	// 	B     bool
	// }
	// var a Vars = Vars{"小米", slice, b}
	// file.Execute(w, a)
	//==============================
	// type Str struct {
	// 	str  string
	// 	str2 string
	// }
	// var str Str
	// str.str = "1904班"
	// str.str2 = "1904班"
	// file.Execute(w, str)
	//==============================
	//file.Execute(w, nil)
}

func post(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //获取表单数据，将数据存入f.Form
	for k, v := range r.Form {
		f.Println(k, "=", v)
	}

}

func main() {
	http.HandleFunc("/gg", index)
	http.HandleFunc("/post", post)
	http.ListenAndServe(":1111", nil)
	/*
		         留言本的数据
		姓名，标题，内容，留言时间，管理员回复
	*/
}
