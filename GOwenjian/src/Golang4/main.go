// Golang4 project main.go
package main

import (
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
	var slice []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var b bool = true
	type Vars struct {
		Name  string
		Slice []int
		B     bool
	}
	var a Vars = Vars{"小米", slice, b}
	file.Execute(w, a)
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
