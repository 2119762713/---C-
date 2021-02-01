// Golang5 project main.go
package main

import (
	"database/sql"
	//"database/sql"
	f "fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func AOL(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("奥里给"))
}

func index(w http.ResponseWriter, r *http.Request) {
	file, err := template.ParseFiles("./html/index.html")
	if err != nil {
		f.Println(err) //不会显示在网页
		w.Write([]byte("404 not found"))
		return
	}
	//==============================
	Db, err2 := sql.Open("mysql", `root:root@tcp(127.0.0.1:3306)/bak?charset=utf8;`)
	if err2 != nil {
		f.Println("数据连接失败", err2)
	}
	res, _ := Db.Query("select id,name from yunwu")

	var a []map[string]string
	for res.Next() {
		var data map[string]string = make(map[string]string)
		var id string
		var name string
		res.Scan(&id, &name)
		data["id"] = id
		data["name"] = name
		a = append(a, data)
	}
	type Vars struct {
		Date []map[string]string
	}
	var aaa Vars = Vars{a}
	file.Execute(w, aaa)
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

func main() {
	http.HandleFunc("/gg", index)
	http.HandleFunc("/AOL", AOL)
	http.ListenAndServe(":5555", nil)
}
