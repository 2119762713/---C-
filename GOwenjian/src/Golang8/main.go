//Golang8
package main

//第一步导入包
import (
	"database/sql"
	"fmt"

	"time"

	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func TimeTo(v string) string {
	loc, _ := time.LoadLocation("Local")
	timeLayout := "2006-01-02 15:04:05"
	tmp, _ := time.ParseInLocation(timeLayout, v, loc)
	timestamp := tmp.Unix()
	res := time.Unix(timestamp, 0).Format(timeLayout)
	return res
}

func index(w http.ResponseWriter, r *http.Request) {
	file, err := template.ParseFiles("./html/index.html")
	if err != nil {
		fmt.Println(err) //不会显示在网页
		w.Write([]byte("404 not found"))
		return
	}

	type Vars struct {
	}
	var aa Vars = Vars{}
	file.Execute(w, aa) //将服务器的响应内容w传给模板
}

func post(w http.ResponseWriter, r *http.Request) {
	//1. 接收表单数据 r.Form是一个map类型
	r.ParseForm()                   //获取表单数据，将数据存入r.Form
	content := r.Form["content"][0] // 从切片取元素出来[]string = [1]
	fmt.Println(content)
	title := r.Form["title"][0]
	username := r.Form["username"][0]
	//通过time包获取时间戳存到表里面
	timeline := int(time.Now().Unix())

	Db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/liuyanben?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}
	//2. SQL语句要自己拼出来
	sql := `insert into liuyanben(username, title, content, timeline)values('` + username + `', 
	'` + title + `', '` + content + `', 
	'` + fmt.Sprintf("%d", timeline) + `'
	)`
	_, err2 := Db.Exec(sql)

	file, err := template.ParseFiles("./html/2.html")
	if err != nil {
		fmt.Println(err) //不会显示在网页
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
	file.Execute(w, aa) //将服务器的响应内容w传给模板
}

//留言列表
func liebiao(w http.ResponseWriter, r *http.Request) {
	file, err := template.ParseFiles("./html/liebiao.html")
	if err != nil {
		fmt.Println(err) //不会显示在网页
		w.Write([]byte("404 not found"))
		return
	}

	//查数据
	Db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/liuyanben?charset=utf8")
	sql := `select * from liuyanben`
	res, err := Db.Query(sql) //QueryRow
	if err != nil {
		fmt.Println(err)
		return
	}

	var all []map[string]string
	for res.Next() {
		var data map[string]string = make(map[string]string)
		var id string
		var username string
		var title string
		var content string
		var timeline string
		var huifu string
		var status string
		res.Scan(&id, &username, &title, &content, &timeline, &huifu, &status)
		data["id"] = id
		data["username"] = username
		data["title"] = title
		data["content"] = content
		data["timeline"] = timeline
		data["huifu"] = huifu
		data["status"] = status
		all = append(all, data)
	}

	type Vars struct {
		Data []map[string]string
	}

	var aa Vars = Vars{all}
	file.Execute(w, aa) //将服务器的响应内容w传给模板

}

func Deletebyid(w http.ResponseWriter, r *http.Request) {
	Db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/liuyanben?charset=utf8")

	id := r.FormValue("id")
	fmt.Println(id)
	sql := `delete from liuyanben where id =` + id + ``
	_, err2 := Db.Exec(sql)
	if err2 == nil {
		fmt.Printf("ID:%v删除成功", id)
	} else {
		fmt.Printf("ID:%v删除失败", id)
	}

	file, err := template.ParseFiles("./html/sanchu.html")
	if err != nil {
		fmt.Println(err) //不会显示在网页
		w.Write([]byte("404 not found"))
		return
	}

	type Vars struct {
		B bool
	}
	var aa Vars
	if err == nil {
		aa.B = true
	} else {
		aa.B = false
	}
	file.Execute(w, aa) //将服务器的响应内容w传给模板
}

func xiugai(w http.ResponseWriter, r *http.Request) {
	file, err := template.ParseFiles("./html/xiugai.html")
	if err != nil {
		fmt.Println(err) //不会显示在网页
		w.Write([]byte("404 not found"))
		return
	}

	type Vars struct {
	}
	var aa Vars = Vars{}
	file.Execute(w, aa) //将服务器的响应内容w传给模板
}

func xiugai2(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	fmt.Println(id)
	Db, _ := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/liuyanben?charset=utf8")
	sql := `select * from liuyanben where id=` + id + ``
	res := Db.QueryRow(sql)
	var data map[string]string = make(map[string]string)
	//var id string
	var username string
	var title string
	var content string
	var timeline string
	var huifu string
	var status string
	res.Scan(&id, &username, &title, &content, &timeline, &huifu, &status)
	data["id"] = id
	data["username"] = username
	data["title"] = title
	data["content"] = content
	data["timeline"] = timeline
	data["huifu"] = huifu
	data["status"] = status
	if username == "" && title == "" && content == "" {
		file, err := template.ParseFiles("./html/xiugai2bucuenzaai.html")
		if err != nil {
			fmt.Println(err) //不会显示在网页
			w.Write([]byte("404 not found"))
			return
		}

		type Vars struct {
		}
		var aa Vars = Vars{}
		file.Execute(w, aa) //将服务器的响应内容w传给模板
	} else {
		//=============================================
		//存在的情况

		file, err := template.ParseFiles("./html/xiugai2.html")
		if err != nil {
			fmt.Println(err) //不会显示在网页
			w.Write([]byte("404 not found"))
			return
		}

		type Vars struct {
			Data map[string]string
		}

		var aa Vars = Vars{data}
		file.Execute(w, aa) //将服务器的响应内容w传给模板
	}
}

func update(w http.ResponseWriter, r *http.Request) {
	//1. 接收表单数据 r.Form是一个map类型
	r.ParseForm()                   //获取表单数据，将数据存入r.Form
	content := r.Form["content"][0] // 从切片取元素出来[]string = [1]
	fmt.Println(content)
	//id := r.FormValue("id")
	id := r.Form["id"][0]
	fmt.Println(id)
	title := r.Form["title"][0]
	fmt.Println(title)
	username := r.Form["username"][0]
	fmt.Println(username)
	Db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/liuyanben?charset=utf8")

	//id := r.FormValue("id")
	//fmt.Println(id)
	sql2 := `update  liuyanben  set username=?,title=?,content =? where id =?;`
	_, err2 := Db.Exec(sql2, username, title, content, id)
	if err2 == nil {
		fmt.Printf("ID:%v修改成功", id)
	} else {
		fmt.Printf("ID:%v修改失败", id)
	}

	file, err := template.ParseFiles("./html/xiugaichengg.html")
	if err != nil {
		fmt.Println(err) //不会显示在网页
		w.Write([]byte("404 not found"))
		return
	}

	type Vars struct {
		B bool
	}
	var aa Vars
	if err == nil {
		aa.B = true
	} else {
		aa.B = false
	}
	file.Execute(w, aa) //将服务器的响应内容w传给模板
}

func main() {
	http.HandleFunc("/index", index)       //表单页
	http.HandleFunc("/post", post)         //表单处理页
	http.HandleFunc("/liebiao", liebiao)   //留言列表
	http.HandleFunc("/delete", Deletebyid) //删除留言列表
	http.HandleFunc("/xiugai", xiugai)     //修改留言列表
	http.HandleFunc("/xiugai2", xiugai2)   //修改留言列表
	http.HandleFunc("/update", update)     //修改留言列表
	http.ListenAndServe(":1111", nil)      //创建端口1111  127.0.0.1:1111
}
