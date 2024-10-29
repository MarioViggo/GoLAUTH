package main

import (
	"database/sql"
	"fmt"
	"golauth/db"
	"golauth/handler"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	d, err := sql.Open("mysql", dataSource()+"?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer d.Close()

	cors := os.Getenv("profile") == "prod"
	app := handler.NewApp(db.NewDB(d), cors)
	time.Sleep(30 * time.Second)
	app.Migrate(d)
	err = app.Serve()
}

func dataSource() string {
	host := "localhost"
	pass := "pass"
	if os.Getenv("profile") == "prod" {
		host = "db"
		pass = os.Getenv("db_pass")
	}
	fmt.Println("[App started]")
	return "authuser:" + pass + "@tcp(" + host + ":3306)/authdb"
}
