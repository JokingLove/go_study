package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main1() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test_db"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("成功连接到MySQL数据库！")

	userId := 1
	getUser(db, userId)

	// createUser
	userId = createUser(db, "李四", 22)

	getUsers(db)

	updateUser(db, userId, "李四2", 33)

	getUser(db, userId)

	// deleteUser(db, userId)
}

func getUsers(db *sql.DB) {
	query := "SELECT id, name, age FROM users"
	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	fmt.Println("所有用户： ")

	for rows.Next() {
		var id int
		var name string
		var age int
		if err := rows.Scan(&id, &name, &age); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d \n", id, name, age)
	}
}

func getUser(db *sql.DB, id int) {
	query := "SELECT id, name , age from users where id = ?"
	row := db.QueryRow(query, id)

	var name string
	var age int

	if err := row.Scan(&id, &name, &age); err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("用户ID %d 不存在\n", id)
		} else {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("ID: %d, Name: %s, Age: %d \n", id, name, age)
	}
}

func createUser(db *sql.DB, name string, age int) int {

	query := "INSERT INTO users(name, age) VALUES(?, ?)"
	result, err := db.Exec(query, name, age)
	if err != nil {
		log.Fatal(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return int(lastInsertId)
}

func updateUser(db *sql.DB, id int, name string, age int) int {
	query := "UPDATE users set name = ?, age = ? where id = ?"
	result, err := db.Exec(query, name, age, id)
	if err != nil {
		log.Fatal(err)
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	return int(affectedRows)
}

func deleteUser(db *sql.DB, id int) int {
	query := "delete from users where id = ?"
	result, err := db.Exec(query, id)
	if err != nil {
		log.Fatal(err)
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	return int(affectedRows)
}
