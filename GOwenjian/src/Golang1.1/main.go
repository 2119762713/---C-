// Golang1.1 project main.go
package main

import (
	"database/sql"
	//f "fmt"
	"net/http"

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
	for res.Next() {
		var id string
		var name string
		res.Scan(&id, &name)
		w.Write([]byte(id))
		w.Write([]byte(name))
	}
}
func main() {
	http.HandleFunc("/1", Hello1)
	http.HandleFunc("/2", Hello2)
	http.HandleFunc("/3", Hello3)
	http.HandleFunc("/all", list)
	http.ListenAndServe(":8000", nil)
}
