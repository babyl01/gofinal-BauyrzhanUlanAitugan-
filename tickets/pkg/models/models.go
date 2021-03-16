package models

type Ticket struct {
	IdUser int
	IdFilm int
}

type Tickets []*Ticket
