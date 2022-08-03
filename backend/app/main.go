package main

import (
  //"fmt"
  "log"
  "errors"
  "github.com/jessevdk/go-flags"
   describe "generator-php-entities/v1/backend/app/db/repository"
   connection "generator-php-entities/v1/backend/app/db"
   entity "generator-php-entities/v1/backend/app/generator"
   jbolt "generator-php-entities/v1/backend/app/store/jbolt"
)

type Options struct {
   DbName string `short:"n" long:"db_name" default:"" description:"DB Name"`
   DbHost string `short:"h" long:"db_host" default:"127.0.0.1" description:"DB Host"`
   DbPort string `short:"p" long:"db_port" default:"3306" description:"DB Port"`
   DbUser string `short:"u" long:"db_user" default:"" description:"DB User"`
   DbPassword string `long:"db_password" default:"" description:"DB Password"`
   DbType string `long:"db_type" default:"mysql" description:"Type of DB"`

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

    if len(opts.DbName) > 0 {
        jbolt.Set(bolt.DB, bucket, "last/DB_NAME", opts.DbName)
    } else {
        jDbName := jbolt.Get(bolt.DB, bucket, "last/DB_NAME")

        if len(jDbName) <= 0 {
            return opts, errors.New("DB name is required")
        }

        opts.DbName = jDbName
    }

    if len(opts.DbHost) > 0 {
        jbolt.Set(bolt.DB, bucket, "last/DB_HOST", opts.DbHost)
    } else {
         jDbHost := jbolt.Get(bolt.DB, bucket, "last/DB_HOST")

        if len(jDbHost) <= 0 {
             return opts, errors.New("DB host is required")
        }

        opts.DbHost = jDbHost
    }

    if len(opts.DbPort) > 0 {
        jbolt.Set(bolt.DB, bucket, "last/DB_PORT", opts.DbPort)
    } else {
        jDbPort := jbolt.Get(bolt.DB, bucket, "last/DB_PORT")

        if len(jDbPort) <= 0 {
            return opts, errors.New("DB port is required")
        }

        opts.DbPort = jDbPort
    }

    if len(opts.DbUser) > 0 {
        jbolt.Set(bolt.DB, bucket, "last/DB_USER", opts.DbUser)
    } else {
        jDbUser := jbolt.Get(bolt.DB, bucket, "last/DB_USER")

        if len(jDbUser) <= 0 {
             return opts, errors.New("DB user is required")
        }

        opts.DbUser = jDbUser
    }

    if len(opts.DbPassword) > 0 {
        jbolt.Set(bolt.DB, bucket, "last/DB_PASSWORD", opts.DbPassword)
    } else {
        jDbPass := jbolt.Get(bolt.DB, bucket, "last/DB_PASSWORD")

        if len(jDbPass) <= 0 {
             return opts, errors.New("DB password is required")
        }

        opts.DbPassword = jDbPass
    }

    if len(opts.DbType) > 0 {
        jbolt.Set(bolt.DB, bucket, "last/DB_TYPE", opts.DbType)
    } else {
        jDbType := jbolt.Get(bolt.DB, bucket, "last/DB_TYPE")

        if len(jDbType) <= 0 {
             return opts, errors.New("DB type is required")
        }

        opts.DbType = jDbType
    }

    return opts, nil
}

func getValueByKey(bolt jbolt.Bolt, opts Options, key string) (string, error) {
    if len(opts.DbType) > 0 {
        jbolt.Set(bolt.DB, bucket, "last/DB_TYPE", opts.DbType)
    } else {
        jDbType := jbolt.Get(bolt.DB, bucket, "last/DB_TYPE")

        if len(jDbType) <= 0 {
             return "", errors.New("DB type is required")
        }

        return jDbType, nil
    }
}