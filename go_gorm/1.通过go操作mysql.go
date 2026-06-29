package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// 链接数据库
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gorm")
	if err != nil {
		log.Fatalf("数据库连接失败 %s", err)
	} else {
		fmt.Println("数据库连接成功")
	}

	defer db.Close()

	// 这里scan的字段要对上
	res, error := db.Query("select id,name,age from users where id = ?", 1)
	if error != nil {
		log.Fatalf("查询数据失败 %s", error)
	}

	for res.Next() {
		var id int
		var name string
		var age int
		res.Scan(&id, &name, &age)
		fmt.Println(id, name, age)

	}

}
