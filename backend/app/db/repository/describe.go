package repository

import (
    "database/sql"
)

type DescribeTable struct {
	Field string `db:"Field"`
	Type string `db:"Type"`
	Null string `db:"Null"`
	Key string `db:"Key"`
	Default sql.NullString `db:"Default"`
	Extra string `db:"Extra"`
}

func Get(connection *sql.DB, table string) ([]DescribeTable, error) {
    var infoTable []DescribeTable
    sqlStatement := "DESCRIBE "+table

    rows, err := connection.Query(sqlStatement)
    if err != nil {
        return nil, err
    }
    for rows.Next() {
        var infoRow DescribeTable
        errRow := rows.Scan(&infoRow.Field, &infoRow.Type, &infoRow.Null, &infoRow.Key,  &infoRow.Default, &infoRow.Extra)
        if errRow != nil {
            return nil, errRow
        }
        infoTable = append(infoTable, infoRow)
    }

    return infoTable, err
}