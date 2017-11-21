// model.go

package main

import (
	"database/sql"
)

type namest struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

func (n *namest) createName(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO names(name) VALUES($1) RETURNING name",
		n.Name).Scan(&n.Name)
	db.Exec("UPDATE names SET count = count + 1 WHERE name=$1",
		n.Name)
	if err != nil {
		return err
	}

	return nil
}

func (n *namest) updateCount(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE names SET count = count + 1 WHERE name=$1",
			n.Name)

	return err
}

func getNames(db *sql.DB, start, count int) ([]namest, error) {
	rows, err := db.Query(
		"SELECT name, count FROM names LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	names := []namest{}

	for rows.Next() {
		var n namest
		if err := rows.Scan(&n.Name, &n.Count); err != nil {
			return nil, err
		}
		names = append(names, n)
	}

	return names, nil
}

func (n *namest) deleteCounts(db *sql.DB) error {
	_, err := db.Exec("UPDATE names SET count=0 WHERE count>0")

	return err
}
