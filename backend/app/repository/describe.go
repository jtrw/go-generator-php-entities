package repository

import (
    "database/sql"
    "log"
)

type DescribeTable struct {
	Field string `db:"Field"`
	Type string `db:"Type"`
	Null string `db:"Null"`
	Key string `db:"Key"`
	Default sql.NullString `db:"Default"`
	Extra string `db:"Extra"`
}

func Get(connection *sql.DB, table string) (DescribeTable, error) {
    var describe DescribeTable

    sqlStatement := `DESCRIBE users`
    rows, err := connection.Query(sqlStatement)
    if err != nil {
        log.Print("d")
        log.Fatal(err)
    }
    for rows.Next() {
        errRow := rows.Scan(&describe.Field, &describe.Type, &describe.Null, &describe.Key,  &describe.Default, &describe.Extra)
        if errRow != nil {
            log.Print("dsss")
            log.Fatal(errRow)
        }
    }


    return describe, err
}