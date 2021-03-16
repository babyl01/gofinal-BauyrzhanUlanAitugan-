package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func (app *App) BuyTicket(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	filmId := r.URL.Query().Get("filmId")

	url := fmt.Sprintf("%s%s", app.apis.tickets, fmt.Sprintf("getTicket?userId=%s&filmId=%s", userId, filmId))

	resp, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	fmt.Fprint(w, string(contents))
}
