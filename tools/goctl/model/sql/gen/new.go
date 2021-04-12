package gen

import (
	"go-zero-study/tools/goctl/model/sql/template"
	"go-zero-study/tools/goctl/util"
)

func genNew(table Table, withCache bool) (string, error) {
	output, err := util.With("new").
		Parse(template.New).
		Execute(map[string]interface{}{
			"withCache":             withCache,
			"upperStartCamelObject": table.Name.ToCamel(),
		})
	if err != nil {
		return "", err
	}
	return output.String(), nil
}
