package parser

import "go-zero-study/tools/goctl/api/spec"

type state interface {
	process(api *spec.ApiSpec) (state, error)
}
