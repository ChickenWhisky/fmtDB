package utils

import "github.com/iancoleman/strcase"


func NormalizeString(input string) string {
	
   return strcase.ToCamel(input)
}
