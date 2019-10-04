package main

import (
	"fmt"
	"flag"

	"github.com/davidfengg/restAPI/database"
	"github.com/davidfengg/restAPI/route"

	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

var err error
// var ipAddress str

func main() {
	ipAddress := flag.String("ip", "localhost:3306", "IP Address")
	flag.Parse()

	// db connection to local mac host
	database.Db, err = sql.Open("mysql", "root:abcd1234@tcp(" + *ipAddress + ")/rest_api")

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connection made")

	// access routing
	route.GetRoutes()

	defer database.Db.Close();
}