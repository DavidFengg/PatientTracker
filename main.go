package main

import (
	"fmt"

	"github.com/davidfengg/restAPI/database"
	"github.com/davidfengg/restAPI/route"

	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

var err error

func main() {
	// db connection to local mac host
	database.Db, err = sql.Open("mysql", "root:abcd1234@tcp(docker.for.mac.localhost:3306)/rest_api")

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connection made")

	// access routing
	route.GetRoutes()

	defer database.Db.Close();
}