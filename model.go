// model.go

package main

import (
	"database/sql"
)

//Primary struct for marshelling and unmarshalling data to a from postgres
type namest struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

//createName - takes name and inserts new row if not found from getName increases count by 1
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

//updateCount - takes getName and if a match is found increases count by 1 
func (n *namest) updateCount(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE names SET count = count + 1 WHERE name=$1",
			n.Name)

	return err
}

//getNames - search and return all names
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

//getName - Scans name table to find match from json input
func (n *namest) getName(db *sql.DB) error {
    return db.QueryRow("SELECT name FROM names WHERE name=$1",
        n.Name).Scan(&n.Name)
}

//deleteCounts - from delete post call delete all counts in count column
func (n *namest) deleteCounts(db *sql.DB) error {
	_, err := db.Exec("UPDATE names SET count=0 WHERE count>0")

	return err
}
