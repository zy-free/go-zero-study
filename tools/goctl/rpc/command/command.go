package command

import (
	"fmt"
	"os"
	"path/filepath"

	"go-zero-study/tools/goctl/rpc/ctx"
	"go-zero-study/tools/goctl/rpc/gen"
	"go-zero-study/tools/goctl/util"
	"github.com/urfave/cli"
)

func Rpc(c *cli.Context) error {
	rpcCtx := ctx.MustCreateRpcContextFromCli(c)
	generator := gen.NewDefaultRpcGenerator(rpcCtx)
	rpcCtx.Must(generator.Generate())
	return nil
}

func RpcTemplate(c *cli.Context) error {
	out := c.String("out")
	idea := c.Bool("idea")
	generator := gen.NewRpcTemplate(out, idea)
	generator.MustGenerate(true)
	return nil
}

func RpcNew(c *cli.Context) error {
	idea := c.Bool("idea")
	arg := c.Args().First()
	if len(arg) == 0 {
		arg = "greet"
	}
	abs, err := filepath.Abs(arg)
	if err != nil {
		return err
	}
	_, err = os.Stat(abs)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		err = util.MkdirIfNotExist(abs)
		if err != nil {
			return err
		}
	}

	dir := filepath.Base(filepath.Clean(abs))

	protoSrc := filepath.Join(abs, fmt.Sprintf("%v.proto", dir))
	templateGenerator := gen.NewRpcTemplate(protoSrc, idea)
	templateGenerator.MustGenerate(false)

	rpcCtx := ctx.MustCreateRpcContext(protoSrc, "", "", idea)
	generator := gen.NewDefaultRpcGenerator(rpcCtx)
	rpcCtx.Must(generator.Generate())
	return nil
}
