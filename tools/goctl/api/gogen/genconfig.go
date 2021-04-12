package gogen

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"go-zero-study/tools/goctl/api/spec"
	"go-zero-study/tools/goctl/api/util"
	"go-zero-study/tools/goctl/vars"
)

const (
	configFile     = "config.go"
	configTemplate = `package config

import {{.authImport}}

type Config struct {
	rest.RestConf
	{{.auth}}
}
`

	jwtTemplate = ` struct {
		AccessSecret string
		AccessExpire int64
	}
`
)

func genConfig(dir string, api *spec.ApiSpec) error {
	fp, created, err := util.MaybeCreateFile(dir, configDir, configFile)
	if err != nil {
		return err
	}
	if !created {
		return nil
	}
	defer fp.Close()

	var authNames = getAuths(api)
	var auths []string
	for _, item := range authNames {
		auths = append(auths, fmt.Sprintf("%s %s", item, jwtTemplate))
	}

	var authImportStr = fmt.Sprintf("\"%s/rest\"", vars.ProjectOpenSourceUrl)
	t := template.Must(template.New("configTemplate").Parse(configTemplate))
	buffer := new(bytes.Buffer)
	err = t.Execute(buffer, map[string]string{
		"authImport": authImportStr,
		"auth":       strings.Join(auths, "\n"),
	})
	if err != nil {
		return nil
	}
	formatCode := formatCode(buffer.String())
	_, err = fp.WriteString(formatCode)
	return err
}
