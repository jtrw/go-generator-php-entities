package main

import (
  "fmt"
  "log"
  "os"
  "github.com/jessevdk/go-flags"
  //"net/http"
  //"io/ioutil"
  //"bytes"
  //"github.com/joho/godotenv"
  "text/template"
   "database/sql"
   _ "github.com/go-sql-driver/mysql"
   describe "generator-php-entities/v1/backend/app/repository"
   "strings"
)

type Options struct {
   DbName string `short:"n" long:"db_name" default:"local_uniquecasino" description:"DB Name"`
   DbHost string `short:"h" long:"db_host" default:"127.0.0.1" description:"DB Host"`
   DbPort string `short:"p" long:"db_port" default:"3306" description:"DB Port"`
   DbUser string `short:"u" long:"db_user" default:"mysql_user" description:"DB User"`
   DbPassword string `long:"db_password" default:"astalavista" description:"DB Password"`
   DbType string `long:"db_type" default:"mysql" description:"Type of DB"`
   Table string `short:"t" long:"table" description:"Table for generate Entity"`
   StoragePath string `short:"s" long:"storage_path" default:"/var/tmp/jtrw_generator_php_entities.db" description:"Storage Path"`
}

func main() {
    var opts Options

    parser := flags.NewParser(&opts, flags.Default)
    _, err := parser.Parse()

    if err != nil {
        log.Fatal(err)
    }

    connection := initDb(opts)
    results, err := describe.Get(connection, opts.Table)

    if err != nil {
        log.Fatal(err)
    }

    printEntity(results);
}

func initDb(opts Options) (*sql.DB) {
    psqlInfo := dsn(opts)
    connection, err := sql.Open(opts.DbType, psqlInfo)
    if err != nil {
        panic(err)
    }

    err = connection.Ping()
    if err != nil {
        panic(err)
    }
    return connection
}

func dsn(opts Options) string {
    return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", opts.DbUser, opts.DbPassword, opts.DbHost, opts.DbPort, opts.DbName)
}

func printEntity(rows []describe.DescribeTable) {
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
        EntityName: "UserEntity",
        Properties: PropertiesData,
        Methods: MethodsData,
    }

    t, err := template.ParseFiles("backend/app/template/entity.gohtml")
    if err != nil {
       panic(err)
    }

    fo, err := os.Create("output.php")
    if err != nil {
        panic(err)
    }
    // close fo on exit and check for its returned error
    defer func() {
        if err := fo.Close(); err != nil {
            panic(err)
        }
    }()

    // Execute the template for each recipient.
    //for _, r := range templateData {
    err = t.Execute(fo, templateData)
    if err != nil {
        log.Println("executing template:", err)
    }
    //}
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
     strings.Contains(t, "date")  {
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