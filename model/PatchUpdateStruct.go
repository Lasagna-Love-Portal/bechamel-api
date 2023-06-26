package model

import (
	"github.com/iancoleman/strcase"
)

type PatchUpdateStruct map[string]interface{}

func (inStruct PatchUpdateStruct) PascalCase() PatchUpdateStruct {
	retStruct := make(PatchUpdateStruct)
	for k, v := range inStruct {
		switch vt := v.(type) {
		case map[string]interface{}:
			retStruct[strcase.ToCamel(k)] = (PatchUpdateStruct)(vt).PascalCase()
		default:
			retStruct[strcase.ToCamel(k)] = vt
		}
	}
	return retStruct
}
