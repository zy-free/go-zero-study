package gen

import (
	"go-zero-study/tools/goctl/model/sql/template"
	"go-zero-study/tools/goctl/util"
)

func genTag(in string) (string, error) {
	if in == "" {
		return in, nil
	}
	output, err := util.With("tag").
		Parse(template.Tag).
		Execute(map[string]interface{}{
			"field": in,
		})
	if err != nil {
		return "", err
	}
	return output.String(), nil
}
