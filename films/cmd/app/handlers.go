package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (app *App) ShowFilm(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.NotFound(w)
		return
	}

	film, err := app.Database.GetFilm(id)
	if err != nil {
		app.ServerError(w, err)
		return
	}
	if film == nil {
		app.NotFound(w)
		return
	}

	json.NewEncoder(w).Encode(film)
}

func (app *App) AddFilm(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	name := r.URL.Query().Get("name")
	price, err := strconv.Atoi(r.URL.Query().Get("price"))

	if err != nil {
		app.ServerError(w, err)
		return
	}

	app.Database.SetFilm(id, price, name)

}
