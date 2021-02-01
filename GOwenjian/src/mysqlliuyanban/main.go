// mysqlliuyanban project main.go
package main

//导入包
import (
	"database/sql"
	f "fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

//创建函数
func Quer(w http.ResponseWriter, r *http.Request) {
	file, err := template.ParseFiles("./moban/liuyanbanmoban.html")
	//导入模板
	if err != nil {
		//模板导入失败
		w.Write([]byte("404 not found"))
		f.Println("模板错误")
		return
	}
	//模板导入成功
	//连接数据库
	Db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/lianxiu?charset=utf8")
	if err != nil {
		w.Write([]byte("数据库连接失败"))
		f.Println("数据库连接失败")
		return
	}
	f.Println("数据库连接成功")
	//查找数据
	sql := `select * from lianxiu` //将所有数据读取出来了
	//查询数据
	res, err := Db.Query(sql)
	if err != nil {
		f.Println("报错,查找数据和查询数据")
		return
	}

	type Data struct {
		Id    int
		Name  string
		Sex   string
		ID    string
		Age   int
		Phone string
	}
	var vas []Data //定义一个结构体切片,传给模板的
	for res.Next() {
		//开始读取数据
		//var []map [string]
		var a = Data{}
		res.Scan(&a.Id, &a.Name, &a.Sex, &a.ID, &a.Age, &a.Phone)
		vas = append(vas, a)
	}
	type Assign struct {
		List []Data
	}
	var b = Assign{vas}
	file.Execute(w, b) //结构体的方式传给模板
}

func delete(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	f.Println(id)
	Db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/lianxiu?charset=utf8")
	if err != nil {
		w.Write([]byte("数据库连接失败"))
		f.Println("数据库连接失败")
		return
	}
	f.Println("数据库连接成功")
	//sql := f.Sprintf("delete from lianxiu where id =")
	_sql := "delete from lianxiu where id ="
	_sql = _sql + f.Sprintf("%s", id)
	_, err = Db.Exec(_sql)
	if err != nil {
		w.Write([]byte("删除失败"))
	} else {
		w.Write([]byte("删除成功"))
	}

}

func main() {
	//路由器设置
	http.HandleFunc("/Quer", Quer)     //首页，进入访问函数Query
	http.HandleFunc("/delete", delete) //留言板，进入访问函数delete
	http.ListenAndServe(":9999", nil)  //端口
}
