package gen

import (
	"strings"

	"go-zero-study/core/collection"
	"go-zero-study/tools/goctl/model/sql/template"
	"go-zero-study/tools/goctl/util"
	"go-zero-study/tools/goctl/util/stringx"
)

func genDelete(table Table, withCache bool) (string, error) {
	keySet := collection.NewSet()
	keyVariableSet := collection.NewSet()
	for fieldName, key := range table.CacheKey {
		if fieldName == table.PrimaryKey.Name.Source() {
			keySet.AddStr(key.KeyExpression)
		} else {
			keySet.AddStr(key.DataKeyExpression)
		}
		keyVariableSet.AddStr(key.Variable)
	}
	var containsIndexCache = false
	for _, item := range table.Fields {
		if item.IsUniqueKey {
			containsIndexCache = true
			break
		}
	}
	camel := table.Name.ToCamel()
	output, err := util.With("delete").
		Parse(template.Delete).
		Execute(map[string]interface{}{
			"upperStartCamelObject":     camel,
			"withCache":                 withCache,
			"containsIndexCache":        containsIndexCache,
			"lowerStartCamelPrimaryKey": stringx.From(table.PrimaryKey.Name.ToCamel()).UnTitle(),
			"dataType":                  table.PrimaryKey.DataType,
			"keys":                      strings.Join(keySet.KeysStr(), "\n"),
			"originalPrimaryKey":        table.PrimaryKey.Name.Source(),
			"keyValues":                 strings.Join(keyVariableSet.KeysStr(), ", "),
		})
	if err != nil {
		return "", err
	}
	return output.String(), nil
}
