package models

import (
	"database/sql"
)

type Database struct {
	*sql.DB
}

func (db *Database) GetTicketList() (Tickets, error) {
	stmt := `SELECT userId, filmId FROM tickets`

	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tickets := Tickets{}

	for rows.Next() {
		r := &Ticket{}

		err := rows.Scan(&r.IdUser, &r.IdUser)
		if err != nil {
			return nil, err
		}

		tickets = append(tickets, r)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tickets, nil
}

func (db *Database) SetTicket(userId, filmId int) (int, error) {
	stmt := `INSERT INTO tickets (userId, filmId) VALUES (?,?)`

	result, err := db.Exec(stmt, userId, filmId)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return int(id), nil
}
