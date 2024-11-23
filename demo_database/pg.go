package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	dsn := "host=localhost port=5432 user=postgres password=postgres dbname=test_db sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	query := "select id, name, age from users"
	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	fmt.Println("用户列表：")
	for rows.Next() {
		var id, age int
		var name string
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("id: %d, Name: %s, Age: %d \n", id, name, age)
	}

	// 插入数据
	insertSQL := "insert into users(name, age) values ($1, $2)"
	_, err = db.Exec(insertSQL, "张三", 19)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("数据插入成功！")
}
