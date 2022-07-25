package main

import (
  //"fmt"
  "log"
  "os"
  "github.com/jessevdk/go-flags"
  //"net/http"
  //"io/ioutil"
  //"bytes"
  //"github.com/joho/godotenv"
  "text/template"
  // "database/sql"
   describe "generator-php-entities/v1/backend/app/repository"
   connection "generator-php-entities/v1/backend/app/db"
   "strings"
)

type Options struct {
   DbName string `short:"n" long:"db_name" default:"local_uniquecasino" description:"DB Name"`
   DbHost string `short:"h" long:"db_host" default:"127.0.0.1" description:"DB Host"`
   DbPort string `short:"p" long:"db_port" default:"3306" description:"DB Port"`
   DbUser string `short:"u" long:"db_user" default:"mysql_user" description:"DB User"`
   DbPassword string `long:"db_password" default:"astalavista" description:"DB Password"`
   DbType string `long:"db_type" default:"mysql" description:"Type of DB"`
   Table string `short:"t" long:"table" required:"true" description:"Table for generate Entity"`
   StoragePath string `short:"s" long:"storage_path" default:"/var/tmp/jtrw_generator_php_entities.db" description:"Storage Path"`
}

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

    generateEntity(getPreparedEntityName(opts.Table), results);
}

func generateEntity(entityName string, rows []describe.DescribeTable) {
    type Property struct {
        Type, Name string
    }

    type Method struct {
        MethodName, TypeMethod, Return string
    }

    type TemplateEntity struct {
        Properties []Property
        Methods []Method
        EntityName string
    }

     var PropertiesData = []Property{}
     var MethodsData = []Method{}

     for _, row := range rows {
         var rowData Property
         var rowMethod Method
        propertyName := getPreparedName(row.Field)
        propertyType := getPreparedType(row.Type)
        rowData.Type = propertyType
        rowData.Name = "$"+propertyName
        PropertiesData = append(PropertiesData, rowData)

        rowMethod.MethodName = "get"+strings.Title(propertyName)
        rowMethod.TypeMethod = propertyType
        rowMethod.Return = propertyName;

        MethodsData = append(MethodsData, rowMethod)

    }

    var templateData = TemplateEntity{
        EntityName: entityName,
        Properties: PropertiesData,
        Methods: MethodsData,
    }

    t, err := template.ParseFiles("backend/app/template/entity.gohtml")
    if err != nil {
       panic(err)
    }

    fo, err := os.Create(entityName+".php")
    if err != nil {
        panic(err)
    }
    // close fo on exit and check for its returned error
    defer func() {
        if err := fo.Close(); err != nil {
            panic(err)
        }
    }()

    err = t.Execute(fo, templateData)
    if err != nil {
        log.Println("executing template:", err)
    }
}

func getPreparedType(t string) (string) {

    if strings.Contains(t, "int") {
        return "int"
    }

    if strings.Contains(t, "decimal") || strings.Contains(t, "float") {
        return "float"
    }

    if strings.Contains(t, "varchar") || strings.Contains(t, "enum") ||
     strings.Contains(t, "time") || strings.Contains(t, "char") ||
     strings.Contains(t, "date") || strings.Contains(t, "text") {
        return "string"
    }

    return t
}

func getPreparedName(name string) (string) {
    words := strings.Split(name, "_")
    name = strings.ToLower(words[0])
    for _, word := range words[1:] {
        name += strings.Title(word)
    }
    return name;
}

func getPreparedEntityName(name string) (string) {
    return strings.Title(name)+"Entity"
}