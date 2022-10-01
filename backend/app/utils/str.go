package utils

import (
   "strings"
)


func SnakeCaseToCamelCase(inputUnderScoreStr string) (camelCase string) {
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

func GetEntityNameFromTableName(name string) (string) {
    nameLen := len(name)
    lastSymbol := string(name[nameLen-1])
    if lastSymbol == "s" {
        name = strings.TrimSuffix(name, "s")
    }

    return SnakeCaseToCamelCase(name)+"Entity"
}