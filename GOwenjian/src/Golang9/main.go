package main

// Golang9 project main.go

import (
	"database/sql"
	f "fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func host(w http.ResponseWriter, r *http.Request) {
	file, err := template.ParseFiles("./html/lianxiu.html")
	if err != nil {
		f.Println(err)
		w.Write([]byte("404 not found"))
	} else {
		file.Execute(w, nil)
	}
}

func add(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //必须解析才行
	stuname := r.PostForm["stuname"][0]
	stusex := r.PostForm["stusex"][0]
	stuage := r.PostForm["stuage"][0]
	stuid := r.PostForm["stuid"][0]
	stuhone := r.PostForm["stuhone"][0]
	sql1 := `insert into lianxiu(stuName,stuSex,stuAge,stuID,stuPhone) values(?,?,?,?,?)`
	Db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/lianxiu?charset=utf8")
	// if err != nil {
	// 	fmt.Println("数据库连接err=", err)
	// }
	file, err := template.ParseFiles("./html/dier.html")
	// if err != nil {
	// 	fmt.Println("解析模板err=", err)
	// }
	type Vas struct {
		B bool
	}
	var b Vas = Vas{false}
	if err != nil {
		b.B = false
		f.Println("数据库连接失败")
	} else {
		_, err := Db.Exec(sql1, stuname, stusex, stuage, stuid, stuhone)
		if err == nil {
			b.B = true
			f.Println("添加成功")
		} else {
			f.Println(err)
		}
	}
	file.Execute(w, b)
}

func liuyanliebiao(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/host", host)
	http.HandleFunc("/add", add)
	http.HandleFunc("/liuyanliebiao", liuyanliebiao)
	http.ListenAndServe(":1200", nil)
}
