package models

type Film struct {
	ID    int
	Name  string
	Price int
}

type Films []*Film
