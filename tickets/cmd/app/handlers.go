package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func (app *App) SetTicket(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(r.URL.Query().Get("userId"))
	filmId, err := strconv.Atoi(r.URL.Query().Get("filmId"))

	rents, err := app.Database.GetTicketList()
	if err != nil {
		app.ServerError(w, err)
		return
	}
	if rents == nil {
		app.NotFound(w)
		return
	}

	for _, r := range rents {
		if r.IdFilm == filmId && r.IdUser == userId {
			fmt.Fprint(w, "already taken")
			return
		}
	}
	app.Database.SetTicket(userId, filmId)
}
