package main

import (
  //"fmt"
  "log"
  "errors"
  "reflect"
  "github.com/jessevdk/go-flags"
  "github.com/oleiade/reflections"
   describe "generator-php-entities/v1/backend/app/db/repository"
   connection "generator-php-entities/v1/backend/app/db"
   entity "generator-php-entities/v1/backend/app/generator"
   jbolt "generator-php-entities/v1/backend/app/store/jbolt"
)

type Options struct {
   DbName string `short:"n" long:"db_name" default:"" description:"DB Name"`
   DbHost string `short:"h" long:"db_host" default:"" description:"DB Host"`
   DbPort string `short:"p" long:"db_port" default:"" description:"DB Port"`
   DbUser string `short:"u" long:"db_user" default:"" description:"DB User"`
   DbPassword string `long:"db_password" default:"" description:"DB Password"`
   DbType string `long:"db_type" default:"" description:"Type of DB"`

   Type string `short:"y" long:"type" default:"entity" description:"Type of generates files"`
   Table string `short:"t" long:"table" required:"true" description:"Table for generate Entity"`
   OutputPath  string `short:"o" long:"output_path" default:"" description:"Path where generation file(s) are saved"`
   StoragePath string `short:"s" long:"storage_path" default:"/var/tmp/jtrw_generator_php_entities.db" description:"Storage Path"`
}

const TYPE_ENTITY string = "entity"
const bucket string = "credentials/db"

func main() {
    var opts Options

    parser := flags.NewParser(&opts, flags.Default)
    _, err := parser.Parse()

    if err != nil {
        log.Fatal(err)
    }

    opts, err = getDbCredentialsFromStore(opts)

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
        var entityOptions = entity.EntityOptions {
            Table: opts.Table,
            OutputPath: opts.OutputPath,
        }
        entity.Generate(entityOptions, results)
    } else {
        log.Fatal("Type of generate files not found")
    }
}

func getDbCredentialsFromStore(opts Options) (Options, error) {
    bolt := jbolt.Open(opts.StoragePath)

    keys := [6]string{"DbPort", "DbName", "DbHost", "DbUser", "DbPassword", "DbType"}
    for _, val := range keys {
        err := opts.fillFromStoreByKey(bolt, val)
        if err != nil {
            return opts, err
        }
    }

    return opts, nil
}

func (opts *Options) fillFromStoreByKey(bolt *jbolt.Bolt, key string) (error) {
    value := opts.GetField(key)
    boltKey := "last/"+key
    if len(value) > 0 {
        jbolt.Set(bolt.DB, bucket, boltKey, value)
    } else {
        value := jbolt.Get(bolt.DB, bucket, boltKey)

        if len(value) <= 0 {
            return errors.New("DB "+key+" is required")
        }

        reflections.SetField(opts, key, value)
    }
    return nil
}

func (opts Options) GetField(field string) string {
    r := reflect.ValueOf(opts)
    f := reflect.Indirect(r).FieldByName(field)
    return string(f.String())
}