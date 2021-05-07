package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Create the database handle, confirm driver is present
	db, err := sql.Open("mysql", "root:Nt*6MDLmn^AjHyeZ@tcp(140.143.143.245:3308)/spider")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rows, _ := db.Query(`SELECT host,user,password From user`)
	defer rows.Close()
	if err != nil {
		fmt.Printf("insert data error: %v\n", err)
		return
	}

	columns, _ := rows.Columns()
	fmt.Println(columns)

	for rows.Next() {

		scanArgs := make([]interface{}, len(columns))
		values := make([]interface{}, len(columns))

		for i := range values {
			scanArgs[i] = &values[i]
		}

		//将数据保存到 record 字典
		err = rows.Scan(scanArgs...)
		for _, col := range values {
			if col != nil {
				fmt.Println(string(col.([]byte)))
			}
		}

	}

	err = rows.Err()
	if err != nil {
		fmt.Printf(err.Error())
	}

}
