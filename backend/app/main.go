package main

import (
  "fmt"
  "log"
  "errors"
  "reflect"
  "github.com/jessevdk/go-flags"
  "github.com/oleiade/reflections"
   describe "generator-php-entities/v1/backend/app/db/repository"
   connection "generator-php-entities/v1/backend/app/db"
   entity "generator-php-entities/v1/backend/app/generator"
   jstore "generator-php-entities/v1/backend/app/store"
   "encoding/json"
)

type Options struct {
   DbName string `short:"n" long:"db_name" default:"" description:"DB Name"`
   DbHost string `short:"h" long:"db_host" default:"" description:"DB Host"`
   DbPort string `short:"p" long:"db_port" default:"" description:"DB Port"`
   DbUser string `short:"u" long:"db_user" default:"" description:"DB User"`
   DbPassword string `long:"db_password" default:"" description:"DB Password"`
   DbType string `long:"db_type" default:"mysql" description:"Type of DB"`

   Type string `short:"y" long:"type" default:"entity" description:"Type of generates files"`
   Table string `short:"t" long:"table" required:"true" description:"Table for generate Entity"`
   OutputPath  string `short:"o" long:"output_path" default:"" description:"Path where generation file(s) are saved"`
   StoragePath string `short:"s" long:"storage_path" default:"/var/tmp/jtrw_generator_php_entities.db" description:"Storage Path"`
   Profile string `long:"profile" description:"Profile's credentials. Command 'list' for display all profiles"`
}

type Profile struct {
    Id int
    Name string
    Type string
}

const TYPE_ENTITY string = "entity"
const bucket string = "credentials/db"
const KEY_LAST_CREDENTIALS string = "last_db_creds"

func main() {
    var opts Options

    parser := flags.NewParser(&opts, flags.Default)
    _, err := parser.Parse()

    if err != nil {
        log.Fatal(err)
    }

    store := jstore.Store {
        StorePath: opts.StoragePath,
    }

    store.JBolt = store.NewStore()

    if len(opts.Profile) > 0 && opts.Profile == "list" {
        log.Println("1 - db_name - mysql, 2 - db_name - pgsql")
        return
    }
    dbSettings := getBdSettings(opts, store)

    conn, cErr := connection.Init(dbSettings)
    if (cErr != nil) {
        log.Fatal(cErr)
    }
    dataSettingsByte, _ := json.Marshal(dbSettings)

    message := jstore.Message {
        Key: KEY_LAST_CREDENTIALS,
        Bucket: bucket,
        DataBite: dataSettingsByte,
    }

    store.Save(&message)

    results, err := describe.Get(conn, opts.Table)

    if err != nil {
        log.Fatal(err)
    }

    if isTypeEntity(opts.Type) {
        var entityOptions = entity.EntityOptions {
            Table: opts.Table,
            OutputPath: opts.OutputPath,
        }
        entity.Generate(entityOptions, results)
    } else {
        log.Fatal("Type of generate files not found")
    }
}

func getBdSettings(opts Options, store jstore.Store) (connection.Settings) {
    dbSettings := connection.Settings{}

    if len(opts.DbHost) > 0 && len(opts.DbPort) > 0 && len(opts.DbUser) > 0 &&
        len(opts.DbPassword) > 0 && len(opts.DbName) > 0 && len(opts.DbType) > 0 {
         dbSettings = connection.Settings {
            Host: opts.DbHost,
            Port: opts.DbPort,
            User: opts.DbUser,
            Pass: opts.DbPassword,
            DbName: opts.DbName,
            Type: opts.DbType,
        }

        return dbSettings
    }

    mess, err := store.Load(bucket, KEY_LAST_CREDENTIALS)

    if err != nil {
        dbSettings = connection.Settings {
            Host: opts.DbHost,
            Port: opts.DbPort,
            User: opts.DbUser,
            Pass: opts.DbPassword,
            DbName: opts.DbName,
            Type: opts.DbType,
        }

        return dbSettings
    }


    json.Unmarshal(mess.DataBite, &dbSettings)

    return dbSettings
}

func isTypeEntity(t string) (bool) {
    return t == TYPE_ENTITY
}

func getDbCredentialsFromStore(opts Options, store jstore.Store) (Options, error) {
    keys := [6]string{"DbPort", "DbName", "DbHost", "DbUser", "DbPassword", "DbType"}
    for _, val := range keys {
        err := opts.fillFromStoreByKey(store, val)
        if err != nil {
            return opts, err
        }
    }

    return opts, nil
}

func (opts *Options) fillFromStoreByKey(store jstore.Store, key string) (error) {

    value := opts.GetField(key)
    boltKey := getBoltKey(key, opts.Profile)

    if len(value) > 0 {
         message := jstore.Message {
            Key: boltKey,
            Bucket: bucket,
            Data: value,
        }

        store.Save(&message)
    } else {
        message, err := store.Load(bucket, boltKey)

        if err != nil {
            return err
        }

        value := message.Data

        if len(value) <= 0 {
            return errors.New("Key:'"+key+"' is required")
        }

        reflections.SetField(opts, key, value)
    }
    return nil
}

func getBoltKey(key, profile string) (string) {
    chunk := "last"

    if len(profile) > 0 {
        chunk =  fmt.Sprintf("profile/%s", profile, key)
    }
    boltKey := fmt.Sprintf("%s/%s", chunk, key)


    return boltKey
}

func (opts Options) GetField(field string) string {
    r := reflect.ValueOf(opts)
    f := reflect.Indirect(r).FieldByName(field)
    return string(f.String())
}