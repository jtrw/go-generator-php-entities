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
)

type Options struct {
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

    printEntity();
}

func printEntity() {
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

     var PropertiesData = []Property{
        {"int", "$id"},
        {"string", "$name"},
    }
    var templateData = TemplateEntity{
        EntityName: "UserEntity",
        Properties: PropertiesData,
        Methods: []Method{
            {"getId", "int", "id"},
            {"getName", "string", "name"},
        },
    }

    t, err := template.ParseFiles("backend/app/template/entity.gohtml")
    if err != nil {
       panic(err)
    }
    //t := template.Must(template.New("template").Parse(tmp))

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