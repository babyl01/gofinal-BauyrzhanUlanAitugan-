package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *App) Routes() http.Handler {
	mux := pat.New()

	mux.Get("/film/id=:id", http.HandlerFunc(app.ShowFilm))
	mux.Get("/film/add", http.HandlerFunc(app.AddFilm))

	return LogRequest(SecureHeaders(mux))
}
