// dating project main.go
package main

import (
	"database/sql"
	"fmt"
	f "fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var B = 0

func post(w http.ResponseWriter, r *http.Request) {
	file, err := template.ParseFiles("./html/moban.html") //导入模板
	if err != nil {                                       //判断err里面有没有数据，如果有模板就导入成功了，如果没有就导入失败了！
		fmt.Println(err) //不会显示在网页
		w.Write([]byte("404 not found"))
		return
	}

	timeline := string(time.Now().Format("2006-01-02 15:04:05"))                       //时间格式
	f.Println(timeline)                                                                //打印时间格式
	b := rand.Intn(10)                                                                 //生成随机数
	f.Println(b)                                                                       //打印生成随机数
	B = b                                                                              //把随机数赋值给全局变量B
	f.Println(B)                                                                       //打印全局变量
	Db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/datinku?charset=utf8") //连接数据库
	if err != nil {                                                                    //判断err里面有没有数据，如果有数据库就连接成功了，如果没有就连接失败了！
		fmt.Println("连接mysql err=", err)
		return
	}
	sql := `select * from test` //将所有数据读取出来了
	res, err := Db.Query(sql)
	if err != nil { //判断err里面有没有数据，如果有读取数据就读取成功了，如果没有就读取失败了！
		f.Println("报错,查找数据和查询数据")
		return
	}

	type Data struct { //结构体
		Id    int
		Name  string
		Sex   string
		ID    string
		Age   string
		Phone string
	}
	var vas []Data   //定义一个结构体切片,传给模板的
	for res.Next() { //for循环res遍历结构体数值
		var a = Data{}
		res.Scan(&a.Id, &a.Name, &a.Sex, &a.ID, &a.Age, &a.Phone)
		vas = append(vas, a)
	}
	// type Assign struct {
	// 	List Data
	// }
	file.Execute(w, vas[b]) //把vas传给模板然后编译模板
}

func hos(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if B == 0 {
		if r.Form["GK"][0] == "2" {
			w.Write([]byte("成功"))
			return
		}
		w.Write([]byte("错误"))
	} else if B == 1 {
		if r.Form["GK"][0] == "2" {
			w.Write([]byte("成功"))
			return
		}
		w.Write([]byte("错误"))
	} else if B == 2 {
		if r.Form["GK"][0] == "2" {
			w.Write([]byte("成功"))
			return
		}
		w.Write([]byte("错误"))
	} else if B == 3 {
		if r.Form["GK"][0] == "2" {
			w.Write([]byte("成功"))
			return
		}
		w.Write([]byte("错误"))
	} else if B == 4 {
		if r.Form["GK"][0] == "2" {
			w.Write([]byte("成功"))
			return
		}
		w.Write([]byte("错误"))
	} else if B == 5 {
		if r.Form["GK"][0] == "2" {
			w.Write([]byte("成功"))
			return
		}
		w.Write([]byte("错误"))
	} else if B == 6 {
		if r.Form["GK"][0] == "2" {
			w.Write([]byte("成功"))
			return
		}
		w.Write([]byte("错误"))
	} else if B == 7 {
		if r.Form["GK"][0] == "2" {
			w.Write([]byte("成功"))
			return
		}
		w.Write([]byte("错误"))
	} else if B == 8 {
		if r.Form["GK"][0] == "2" {
			w.Write([]byte("成功"))
			return
		}
		w.Write([]byte("错误"))
	} else if B == 9 {
		if r.Form["GK"][0] == "2" {
			w.Write([]byte("成功"))
			return
		}
		w.Write([]byte("错误"))
	}
}

func main() {
	http.HandleFunc("/hos", hos)
	http.HandleFunc("/post", post)
	http.ListenAndServe(":9999", nil)
}
