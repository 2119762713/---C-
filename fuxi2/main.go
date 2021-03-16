// fuxi2 project main.go
package main

import (
	"database/sql"
	f "fmt"
	"html/template"
	"net/http"

	//"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func post(w http.ResponseWriter, r *http.Request) {
	//导入模板
	file, err := template.ParseFiles("./html/moban.html")
	if err != nil {
		f.Println("17行模板导入错误\n")
		return
	}

	ids := r.FormValue("id")
	if ids == "" {
		ids = "0"
	}

	Db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/yun?charset=utf8")
	if err != nil {
		f.Println("27行连接数据库失败\n")
		return
	}
	res := Db.QueryRow(`select * from wenti where id > ? limit 1`, ids)

	var Id int
	var WenTing string
	var A string
	var B string
	var C string
	var D string
	var DaAn int
	res.Scan(&Id, &WenTing, &A, &B, &C, &D, &DaAn)
	type Alis struct {
		Id      int
		WenTing string
		A       string
		B       string
		C       string
		D       string
		DaAn    int
	}
	var aa Alis = Alis{Id, WenTing, A, B, C, D, DaAn}

	file.Execute(w, aa)
}

func Ho(w http.ResponseWriter, r *http.Request) {
	//导入模板
	file, err := template.ParseFiles("./html/moban1.html")
	if err != nil {
		f.Println("Ho模板导入错误\n")
		return
	}
	r.ParseForm()
	xx := r.FormValue("xx")
	id := r.FormValue("id")
	f.Println("数据库答案=", xx)
	f.Println("题目的id=", id)

	Db, _ := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/yun?charset=utf8")
	res := Db.QueryRow(`select * from wenti where id =?`, id) //读取数据库数据
	var Id int
	var WenTing string
	var A string
	var B string
	var C string
	var D string
	var DaAn int
	res.Scan(&Id, &WenTing, &A, &B, &C, &D, &DaAn)

	daAn := f.Sprintf("%d", DaAn) //转换数据类型

	type Res struct {
		Tip  string
		Id   string
		Isok bool
	}

	var a = Res{}
	a.Id = id

	if xx == daAn {
		a.Tip = "恭喜您，回答正确"
		a.Isok = true
	} else {
		a.Tip = "很遗憾，回答错误"
		if xx == "1" {
			a.Tip = a.Tip + ",正确答案是" + A
		} else if daAn == "2" {
			a.Tip = a.Tip + ",正确答案是" + B
		} else if daAn == "3" {
			a.Tip = a.Tip + ",正确答案是" + C
		} else {
			a.Tip = a.Tip + ",正确答案是" + D
		}
		a.Isok = false
		//session  会话机制

	}
	file.Execute(w, a)
}

func main() {
	http.HandleFunc("/post", post)    //创建路由
	http.HandleFunc("/Ho", Ho)        //创建路由
	http.ListenAndServe(":9999", nil) //创建端口
}
