package main

import (
  //"fmt"
  "log"
  //"os"
  "github.com/jessevdk/go-flags"
  //"net/http"
  //"io/ioutil"
  //"bytes"
  //"github.com/joho/godotenv"
  // "database/sql"
   describe "generator-php-entities/v1/backend/app/db/repository"
   connection "generator-php-entities/v1/backend/app/db"
   entity "generator-php-entities/v1/backend/app/generator"
   "strings"
)

type Options struct {
   DbName string `short:"n" long:"db_name" default:"local_uniquecasino" description:"DB Name"`
   DbHost string `short:"h" long:"db_host" default:"127.0.0.1" description:"DB Host"`
   DbPort string `short:"p" long:"db_port" default:"3306" description:"DB Port"`
   DbUser string `short:"u" long:"db_user" default:"mysql_user" description:"DB User"`
   DbPassword string `long:"db_password" default:"astalavista" description:"DB Password"`
   DbType string `long:"db_type" default:"mysql" description:"Type of DB"`

   Type string `short:"y" long:"type" default:"entity" description:"Type of generates files"`
   Table string `short:"t" long:"table" required:"true" description:"Table for generate Entity"`
   StoragePath string `short:"s" long:"storage_path" default:"/var/tmp/jtrw_generator_php_entities.db" description:"Storage Path"`
}

const TYPE_ENTITY string = "entity"

func main() {
    var opts Options

    parser := flags.NewParser(&opts, flags.Default)
    _, err := parser.Parse()

    if err != nil {
        log.Fatal(err)
    }

    dbSettings := connection.Settings {
        Host: opts.DbHost,
        Port: opts.DbPort,
        User: opts.DbUser,
        Pass: opts.DbPassword,
        DbName: opts.DbName,
        Type: opts.DbType,
    }

    conn, cErr := connection.Init(dbSettings)
    if (cErr != nil) {
        log.Fatal(cErr)
    }
    results, err := describe.Get(conn, opts.Table)

    if err != nil {
        log.Fatal(err)
    }
    if opts.Type == TYPE_ENTITY {
        entity.Generate(getPreparedEntityName(opts.Table), results)
    } else {
        log.Fatal("Type of generate files not found")
    }

}


func getPreparedEntityName(name string) (string) {
    return strings.Title(name)+"Entity"
}