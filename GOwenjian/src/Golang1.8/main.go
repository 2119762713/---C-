// Golang1.8 project main.go
package main

import (
	"database/sql"
	//f "fmt"
	"net/http"
	//"html/template"
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

func list(w http.ResponseWriter, r *http.Request) {
	Db, _ := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/bak?charset=utf8;")
	res, _ := Db.Query("select id,name from yunwu")
	w.Header().Set("content-Type", "text/html") //响应头
	strings := `<table border="1" cellpadding="0" style="border:1px solid #ccc">`
	//str := `<table border="1" cellpadding="0" style="border:1px solid #ccc">`
	for res.Next() {
		var id string
		var name string
		res.Scan(&id, &name)
		strings = strings + `<tr>
		<td>` + id + `</td> 
		<td>` + name + `</td> 
		</tr>`
		w.Write([]byte(id))
		w.Write([]byte(name))
		//w.Write([]byte("<&gthr/>"))
	}
	w.Write([]byte(strings))
	// str = str + `<tr> <td>4</td> <td>云雾</td> </tr></table>` //字符串拼接
	// w.Write([]byte(str))
	/* str := <table border="1" cellpadding="0" style="border:1px solid #ccc"><tr> <td>4</td> <td>云雾</td> </tr></table>
	w.Write([]byte(str))
	*/
}

func main() {
	http.HandleFunc("/1", Hello1)
	http.HandleFunc("/2", Hello2)
	http.HandleFunc("/3", Hello3)
	http.HandleFunc("/allo", list)
	http.ListenAndServe(":8000", nil)
}
