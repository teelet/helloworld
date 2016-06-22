package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
	"strconv"
)

type item struct {
	id int
	name string
}

func main()  {
	fmt.Println("helloworld")
	db, err := sql.Open("mysql", "test:123456@tcp(www.teelet.com:3306)/test?charset=utf8")
	defer db.Close()
	checkErr(err)
	err = db.Ping()
	checkErr(err)
	stmt, err := db.Prepare("select * from test limit 10")
	checkErr(err)
	rows, err := stmt.Query()
	checkErr(err)
	for rows.Next() {
		i := new(item)
		rows.Scan(&(i.id), &(i.name))
		fmt.Println(strconv.Itoa(i.id) + " : " + i.name)
	}
}

func checkErr(err error){
	if err != nil {
		log.Panic(err.Error())
	}
}