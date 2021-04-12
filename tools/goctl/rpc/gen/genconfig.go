package gen

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"go-zero-study/tools/goctl/util"
)

const configTemplate = `package config

import "go-zero-study/zrpc"

type Config struct {
	zrpc.RpcServerConf
}
`

func (g *defaultRpcGenerator) genConfig() error {
	configPath := g.dirM[dirConfig]
	fileName := filepath.Join(configPath, fileConfig)
	if util.FileExists(fileName) {
		return nil
	}
	return ioutil.WriteFile(fileName, []byte(configTemplate), os.ModePerm)
}
