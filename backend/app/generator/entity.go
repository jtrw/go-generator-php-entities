package generator

import (
  "fmt"
  "log"
  "os"
  "text/template"
   "strings"
)

type Property struct {
    Type, Name string
}

type Method struct {
    MethodName, TypeMethod, Return, OriginName string
}

type Use struct {
    Name string
}

type TemplateEntity struct {
    Uses []Use
    Properties []Property
    Methods []Method
    EntityName string
}

type EntityOptions struct {
    Name, OutputPath, Template string
}

func Generate(opts EntityOptions, rows []Info) {
    entityName := opts.Name
    fmt.Printf("Enter Entity name default '%s': ", entityName)
    fmt.Scanln(&entityName)

    var PropertiesData = []Property{}
    var MethodsData = []Method{}
    var usesData = []Use{}

    isUseDatetime := false

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
        rowMethod.Return = propertyName
        rowMethod.OriginName = row.Field
        MethodsData = append(MethodsData, rowMethod)

        if isTime(row.Type) {
            isUseDatetime = true
        }
    }
    if isUseDatetime {
        var useData Use
        useData.Name = "DateTime"
        usesData = append(usesData, useData)
    }

    var templateData = TemplateEntity{
        Uses: usesData,
        EntityName: entityName,
        Properties: PropertiesData,
        Methods: MethodsData,
    }

    t, err := template.ParseFiles(opts.Template)
    if err != nil {
       panic(err)
    }
    oFile := entityName+".php"
    oPath := ""
    if len(opts.OutputPath) > 0 {
        oPath = opts.OutputPath
    }

    fo, err := os.Create(oPath+oFile)
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
    if isInt(t) {
        return "int"
    }

    if isFloat(t) {
        return "float"
    }

    if isString(t) {
        return "string"
    }

    if isTime(t) {
        return "DateTime"
    }

    if isArray(t) {
        return "array"
    }

    log.Println("Undefined Type: "+t);
    return t
}

func isInt(t string) (bool) {
    return t == "int" || strings.Contains(t, "int")
}

func isFloat(t string) (bool) {
    return t == "string" || strings.Contains(t, "decimal") || strings.Contains(t, "float")
}

func isString(t string) (bool) {
    return t == "bool" || strings.Contains(t, "varchar") || strings.Contains(t, "enum") ||
        strings.Contains(t, "char") ||strings.Contains(t, "text")
}

func isTime(t string) (bool) {
    return t == "DateTime" || strings.Contains(t, "time") || strings.Contains(t, "date")
}

func isArray(t string) (bool) {
    return t == "array"
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