// qimo project main.go
package main

import (
	"database/sql"
	"fmt"
	f "fmt"
	"html/template"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func pos(w http.ResponseWriter, r *http.Request) {
	file, err := template.ParseFiles("./moban/qimobang.html")
	if err != nil {
		w.Write([]byte("404 not found"))
		f.Println("模板错误")
		return
	}
	file.Execute(w, nil)
}

func post(w http.ResponseWriter, r *http.Request) {
	Db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/lianxiu?charset=utf8")
	if err != nil {
		w.Write([]byte("数据库连接失败"))
		f.Println("数据库连接失败")
		return
	}
	f.Println("数据库连接成功")

	r.ParseForm()
	name := r.Form["name"][0]
	f.Println(name)
	jiage := r.Form["jiage"][0]
	pinpai := r.Form["pinpai"][0]
	shulian := r.Form["shulian"][0]
	shijian := time.Now().Format("2006-01-02 15:04:05")
	miaoshu := r.Form["miaoshu"][0]
	//shijian := int(time.Now().Unix())

	sqlaolg := `insert into kkj(name,jiage,pinpai,shulian,shijian,miaoshu) values(?,?,?,?,?,?)`
	_, err2 := Db.Exec(sqlaolg, name, jiage, pinpai, shulian, shijian, miaoshu)
	fmt.Print(err2)
	if err == nil {
		f.Println("添加成功")
	}
	file, err := template.ParseFiles("./moban/tianjia.html")
	if err != nil {
		f.Println(err)
		w.Write([]byte("404 not found"))
		return
	}

	type Vars struct {
		B bool
	}
	var aa Vars
	if err2 == nil {
		aa.B = true
	} else {
		aa.B = false
	}
	file.Execute(w, aa)
}

func aolg(w http.ResponseWriter, r *http.Request) {
	file, err := template.ParseFiles("./moban/biaoge.html")
	if err != nil {
		w.Write([]byte("404 not found"))
		f.Println("模板错误")
		return
	}
	Db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/lianxiu?charset=utf8")
	if err != nil {
		w.Write([]byte("数据库连接失败"))
		f.Println("数据库连接失败")
		return
	}
	sql := `select * from kkj`

	res, err := Db.Query(sql)
	if err != nil {
		f.Println("报错,查找数据和查询数据")
		return
	}

	type Data struct {
		Id      int
		Name    string
		Jiage   int
		Pinpai  string
		Shulian int
		Shijian string
		Miaoshu string
	}
	var YUN []Data
	for res.Next() {

		var a = Data{}
		res.Scan(&a.Id, &a.Name, &a.Jiage, &a.Pinpai, &a.Shulian, &a.Shijian, &a.Miaoshu)
		YUN = append(YUN, a)
	}
	type Yaolg struct {
		Qq []Data
	}
	var ao = Yaolg{YUN}
	file.Execute(w, ao)
}

func main() {
	http.HandleFunc("/pos", pos)
	http.HandleFunc("/post", post)
	http.HandleFunc("/aolg", aolg)
	http.ListenAndServe(":1201", nil)
}
