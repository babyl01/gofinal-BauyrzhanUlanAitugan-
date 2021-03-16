package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

type apis struct {
	films   string
	users   string
	tickets string
}

type App struct {
	apis apis
}

var token string

func main() {
	addr := flag.String("addr", ":8000", "HTTP network address")
	filmsAPI := flag.String("filmsAPI", "http://localhost:8001/film/", "Films API")
	usersAPI := flag.String("usersAPI", "http://localhost:8002/user/", "Users API")
	ticketsAPI := flag.String("ticketsAPI", "http://localhost:8003/", "Tickets API")
	flag.Parse()

	app := &App{
		apis: apis{
			films:   *filmsAPI,
			users:   *usersAPI,
			tickets: *ticketsAPI,
		},
	}

	fmt.Println("enter token")
	fmt.Scanln(&token)

	log.Printf("Server listening on %s", *addr)
	err := http.ListenAndServe(*addr, app.Routes())
	log.Fatal(err)
}
