package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *App) Routes() http.Handler {
	mux := pat.New()

	mux.Get("/getTicket", http.HandlerFunc(app.SetTicket))

	return LogRequest(SecureHeaders(mux))
}
