// Golang7 project main.go
package main

import (
	"database/sql"
	f "fmt"
	"html/template"
	"net/http"
	"time"

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

}

func post(w http.ResponseWriter, r *http.Request) {
	//====================================
	r.ParseForm() //获取表单数据，将数据存入f.Form
	content := r.Form["content"][0]
	title := r.Form["title"][0]
	username := r.Form["username"][0]
	timeline := int(time.Now().Unix())
	//====================================
	Db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/liuyanben?charset=utf8")
	if err != nil {
		f.Println("数据连接失败", err)
		return
	}
	//===============================
	sql := ` insert into liuyanben(username,title,content,timeline)values('` + username + `',
	'` + title + `',
	'` + content + `',
	'` + f.Sprintf("%v", timeline) + `')`

	_, err3 := Db.Exec(sql)
	if err3 == nil {
		f.Println("添加成功")
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
