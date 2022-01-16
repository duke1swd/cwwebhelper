package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "yaws"
	password = "159265"
	dbname   = "chapelwood"
)

var (
	psqlconn string
)

func init() {
	psqlconn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
}

func usersCmd() {

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")

	// select email, listsrv, wwwsrv, burpsrv from email;
	rows, err := db.Query("select email, listsrv, wwwsrv, burpsrv from email")
	CheckError(err)
	defer rows.Close()

	for rows.Next() {
		var (
			email string
			listsrv bool
			wwwsrv bool
			burpsrv bool
		)

		err = rows.Scan(&email, &listsrv, &wwwsrv, &burpsrv)
		CheckError(err)
		fmt.Printf("%s,%v,%v,%v\n", email, listsrv, wwwsrv, burpsrv)
	}
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
