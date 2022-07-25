package db

import (
  "fmt"
   "database/sql"
   _ "github.com/go-sql-driver/mysql"
   "errors"
)

type Settings struct {
    Host string
    Port string
    User string
    Pass string
    DbName string
    Type string
	Connection *sql.DB
}

func Init(s Settings) (*sql.DB, error) {
    var dsn string

    switch s.Type {
        case "mysql":
           dsn = getMysqlDsn(s)
        case "pgsql":
            dsn = getPgsqlDsn(s)
        default:
            errors.New("DB type is not supported")
    }

    connection, err := sql.Open(s.Type, dsn)
    if err != nil {
        return nil, err
    }

    err = connection.Ping()
    if err != nil {
        return nil, err
    }
    s.Connection = connection
    return connection, nil
}

func getMysqlDsn(s Settings) (string) {
    return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", s.User, s.Pass, s.Host, s.Port, s.DbName)
}

func getPgsqlDsn(s Settings) (string) {
    return fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        s.Host,
        s.Port,
        s.User,
        s.Pass,
        s.DbName)
}
