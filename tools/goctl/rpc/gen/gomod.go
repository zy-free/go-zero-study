package gen

import (
	"fmt"

	"go-zero-study/core/logx"
	"go-zero-study/tools/goctl/rpc/execx"
)

func (g *defaultRpcGenerator) initGoMod() error {
	if !g.Ctx.IsInGoEnv {
		projectDir := g.dirM[dirTarget]
		cmd := fmt.Sprintf("go mod init %s", g.Ctx.ProjectName.Source())
		output, err := execx.Run(fmt.Sprintf(cmd), projectDir)
		if err != nil {
			logx.Error(err)
			return err
		}
		g.Ctx.Info(output)
	}
	return nil
}
