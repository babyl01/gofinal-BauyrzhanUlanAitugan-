package main

import (
	"awesomeProject2/tickets/pkg/models"
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

type App struct {
	Database *models.Database
}

func main() {
	addr := flag.String("addr", ":8003", "HTTP network address")
	dsn := flag.String("dsn", "root:root@/cinema?parseTime=true", "MySQL DSN")
	flag.Parse()

	db := connect(*dsn)
	defer db.Close()

	app := &App{
		Database: &models.Database{DB: db},
	}

	log.Printf("Renting service listening on %s", *addr)
	err := http.ListenAndServe(*addr, app.Routes())
	log.Fatal(err)

}

func connect(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
