package db

import (
  "fmt"
  "log"
  "text/template"
   "database/sql"
   _ "github.com/go-sql-driver/mysql"
   "errors"
)

func init(opts Options) (*sql.DB, error) {
    var dsn string

    switch opts.DbType {
        case "mysql":
           dsn = getMysqlDsn(opts)
        case "pgsql":
            dsn = getPgsqlDsn(opts)
        default:
            errors.New("DB type is not supported")
    }

    connection, err := sql.Open(opts.DbType, dsn)
    if err != nil {
        return nil, err
    }

    err = connection.Ping()
    if err != nil {
        return nil, err
    }

    return connection, nil
}

func getMysqlDsn(opts Options) (string) {
    return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", opts.DbUser, opts.DbPassword, opts.DbHost, opts.DbPort, opts.DbName)
}

func getPgsqlDsn(opts Options) (string) {
    return fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        opts.DbHost,
        opts.DbPort,
        opts.DbUser,
        opts.DbPassword,
        opts.DbName)
}
