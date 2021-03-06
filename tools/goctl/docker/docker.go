package docker

import (
	"errors"

	"go-zero-study/tools/goctl/gen"
	"github.com/urfave/cli"
)

func DockerCommand(c *cli.Context) error {
	goFile := c.String("go")
	if len(goFile) == 0 {
		return errors.New("-go can't be empty")
	}

	return gen.GenerateDockerfile(goFile, "-f", "etc/config.yaml")
}
