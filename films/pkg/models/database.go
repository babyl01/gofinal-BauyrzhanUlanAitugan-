package models

import (
	"database/sql"
)

type Database struct {
	*sql.DB
}

func (db *Database) GetFilm(id int) (*Film, error) {
	stmt := `SELECT id, name, price FROM films
	WHERE id = ?`

	row := db.QueryRow(stmt, id)

	film := &Film{}

	err := row.Scan(&film.ID, &film.Name, &film.Price)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return film, nil
}

func (db *Database) SetFilm(id int, price int, name string) (int, error) {
	stmt := `INSERT INTO films (id, name, price) VALUES (?,?,?)`

	result, err := db.Exec(stmt, id, name, price)
	if err != nil {
		return 0, err
	}

	idNew, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return int(idNew), nil
}
