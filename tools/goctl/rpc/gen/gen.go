package gen

import (
	"go-zero-study/tools/goctl/rpc/ctx"
	"go-zero-study/tools/goctl/rpc/parser"
)

const (
	dirTarget          = "dirTarget"
	dirConfig          = "config"
	dirEtc             = "etc"
	dirSvc             = "svc"
	dirServer          = "server"
	dirLogic           = "logic"
	dirPb              = "pb"
	dirInternal        = "internal"
	fileConfig         = "config.go"
	fileServiceContext = "servicecontext.go"
)

type defaultRpcGenerator struct {
	dirM map[string]string
	Ctx  *ctx.RpcContext
	ast  *parser.PbAst
}

func NewDefaultRpcGenerator(ctx *ctx.RpcContext) *defaultRpcGenerator {
	return &defaultRpcGenerator{
		Ctx: ctx,
	}
}

func (g *defaultRpcGenerator) Generate() (err error) {
	g.Ctx.Info("generating code...")
	defer func() {
		if err == nil {
			g.Ctx.Success("Done.")
		}
	}()
	err = g.createDir()
	if err != nil {
		return
	}

	err = g.initGoMod()
	if err != nil {
		return
	}

	err = g.genEtc()
	if err != nil {
		return
	}

	err = g.genPb()
	if err != nil {
		return
	}

	err = g.genConfig()
	if err != nil {
		return
	}

	err = g.genSvc()
	if err != nil {
		return
	}

	err = g.genLogic()
	if err != nil {
		return
	}

	err = g.genHandler()
	if err != nil {
		return
	}

	err = g.genMain()
	if err != nil {
		return
	}

	err = g.genCall()
	if err != nil {
		return
	}

	return
}
