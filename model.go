// model.go

package main

import (
    "database/sql"
)

type namest struct {
    Name  string  `json:"name"`
    Count int     `json:"count"`
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
