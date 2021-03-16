package models

type User struct {
	ID      int
	Name    string
	Surname string
}

type Users []*User
