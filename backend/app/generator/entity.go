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
    MethodName, TypeMethod, Return string
}

type TemplateEntity struct {
    Properties []Property
    Methods []Method
    EntityName string
}

type EntityOptions struct {
    Table, OutputPath string
}

func Generate(opts EntityOptions, rows []Info) {
    tableName := opts.Table
    entityName := getEntityNameFromTableName(tableName)
    fmt.Printf("Enter Entity name default '%s': ", entityName)
    fmt.Scanln(&entityName)

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

func getEntityNameFromTableName(name string) (string) {
    nameLen := len(name)
    lastSymbol := string(name[nameLen-1])
    if lastSymbol == "s" {
        name = strings.TrimSuffix(name, "s")
    }

    return snakeCaseToCamelCase(name)+"Entity"
}

func snakeCaseToCamelCase(inputUnderScoreStr string) (camelCase string) {
    isToUpper := false
    for k, v := range inputUnderScoreStr {
        if k == 0 {
            camelCase = strings.ToUpper(string(inputUnderScoreStr[0]))
        } else {
            if isToUpper {
                camelCase += strings.ToUpper(string(v))
                isToUpper = false
            } else {
                 if v == '_' {
                    isToUpper = true
                 } else {
                    camelCase += string(v)
                 }
            }
        }
    }
    return
}