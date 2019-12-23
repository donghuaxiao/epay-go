package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"reflect"
)

/*
type struct User {
    Id int64
    Name string
    Password string
    Email string
    CreateTime time.Time
    LoginTimes int
    Status int
}
*/

func printResult(query *sql.Rows) {
	column, _ := query.Columns()

	values := make([][]byte, len(column))
	scans := make([]interface{}, len(column))
	for i := range values {
		scans[i] = &values[i]
	}

	results := make(map[int]map[string]string)
	i := 0

	for query.Next() {
		if err := query.Scan(scans...); err != nil {
			fmt.Println(err)
			return
		}

		row := make(map[string]string)

		for k, v := range values {
			key := column[k]
			row[key] = string(v)
		}

		results[i] = row
		i++
	}

	for k, v := range results {
		fmt.Println(k, v)
	}
}

func count(db *sql.DB, countstr string, args ...interface{}) int {
	query := db.QueryRow(countstr, args...)
	var count int
	err := query.Scan(&count)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("no row")
	case err != nil:
		log.Fatalf("query error:%v\n", err)
	default:
		log.Printf("count : %d\n", count)
	}
	return count
}

func main() {

	db, err := sql.Open("mysql", "root:Passw0rd@tcp(127.0.0.1:3306)/test?charset=utf8")

	fmt.Println(reflect.TypeOf(db))

	if err != nil {
		log.Fatal(err)
	}
	query, err := db.Query("select * from t_users")

	if err != nil {
		log.Fatal(err)
	}

	v := reflect.ValueOf(query)
	fmt.Println(v)

	printResult(query)

	count := count(db, "select count(*) from test.t_users")
	fmt.Println(count)
}
