// pub project main.go
package main{
	
}
//用来初始化一些值
var Db *sql.DB //全局变量

func init() {
	Db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/liuyanben?charset=utf8")
	if err != nil {
		f.Println(err)
		return
	}
}