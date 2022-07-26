package repository

import (
    "database/sql"
     field "generator-php-entities/v1/backend/app/generator"
)

type DescribeTable struct {
	Field string `db:"Field"`
	Type string `db:"Type"`
	Null string `db:"Null"`
	Key string `db:"Key"`
	Default sql.NullString `db:"Default"`
	Extra string `db:"Extra"`
}

func Get(connection *sql.DB, table string) ([]field.Info, error) {
    var infoTable []field.Info
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
        filedInfo := field.Info {
            Field: infoRow.Field,
            Type: infoRow.Type,
        }
        infoTable = append(infoTable, filedInfo)
    }

    return infoTable, err
}