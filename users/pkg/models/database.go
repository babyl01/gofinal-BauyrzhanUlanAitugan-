package models

import (
	"database/sql"
)

type Database struct {
	*sql.DB
}

func (db *Database) GetUser(id int) (*User, error) {
	stmt := `SELECT id, name, surname FROM user
	WHERE id = ?`

	row := db.QueryRow(stmt, id)

	u := &User{}

	err := row.Scan(&u.ID, &u.Name, &u.Surname)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return u, nil
}
