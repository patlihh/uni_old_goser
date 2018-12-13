package main

import (
	"database/sql"
	"fmt"
	//	"github.com/go-sql-driver/mysql"
	"log"
	"upks/unet"
)

func QueryUnDb() {

	db, err := sql.Open("mysql", "root:unitone@/ips")
	if err != nil {
		log.Fatalf("Open database error: %s\n", err)
	}
	defer db.Close()

	fmt.Println("ping0")

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ping1")

	//	insert(db)

	rows, err := db.Query("select * from sys_user")
	if err != nil {
		log.Println(err)
	}

	fmt.Println("query1")

	defer rows.Close()
	var id, name, pwd, state string

	for rows.Next() {

		fmt.Println("row0")

		err := rows.Scan(&id, &name, &pwd, &state)

		fmt.Println("row1")

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("row2")
		log.Println(id, name)

	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

//initial listener and run
func main() {

	fmt.Printf("ups running ...\n")

	unet.TcpManage()

}
