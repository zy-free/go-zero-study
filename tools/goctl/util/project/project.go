package project

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"go-zero-study/tools/goctl/rpc/execx"
)

const (
	constGo          = "go"
	constProtoC      = "protoc"
	constGoMod       = "go env GOMOD"
	constGoPath      = "go env GOPATH"
	constProtoCGenGo = "protoc-gen-go"
)

type (
	Project struct {
		Path    string // Project path name
		Name    string // Project name
		Package string // The service related package
		// true-> project in go path or project init with go mod,or else->false
		IsInGoEnv bool
		GoMod     GoMod
	}

	GoMod struct {
		Module string // The gomod module name
		Path   string // The gomod related path
	}
)

func Prepare(projectDir string, checkGrpcEnv bool) (*Project, error) {
	_, err := exec.LookPath(constGo)
	if err != nil {
		return nil, err
	}

	if checkGrpcEnv {
		_, err = exec.LookPath(constProtoC)
		if err != nil {
			return nil, err
		}

		_, err = exec.LookPath(constProtoCGenGo)
		if err != nil {
			return nil, err
		}
	}

	var (
		goMod, module string
		goPath        string
		name, path    string
		pkg           string
	)

	ret, err := execx.Run(constGoMod, projectDir)
	if err != nil {
		return nil, err
	}

	goMod = strings.TrimSpace(ret)
	if goMod == os.DevNull {
		goMod = ""
	}

	ret, err = execx.Run(constGoPath, "")
	if err != nil {
		return nil, err
	}

	goPath = strings.TrimSpace(ret)
	src := filepath.Join(goPath, "src")
	var isInGoEnv = true
	if len(goMod) > 0 {
		path = filepath.Dir(goMod)
		name = filepath.Base(path)
		data, err := ioutil.ReadFile(goMod)
		if err != nil {
			return nil, err
		}

		module, err = matchModule(data)
		if err != nil {
			return nil, err
		}
	} else {
		pwd, err := filepath.Abs(projectDir)
		if err != nil {
			return nil, err
		}

		if !strings.HasPrefix(pwd, src) {
			name = filepath.Clean(filepath.Base(pwd))
			path = projectDir
			pkg = name
			isInGoEnv = false
		} else {
			r := strings.TrimPrefix(pwd, src+string(filepath.Separator))
			name = filepath.Dir(r)
			if name == "." {
				name = r
			}
			path = filepath.Join(src, name)
			pkg = r
		}
		module = name
	}

	return &Project{
		Name:      name,
		Path:      path,
		Package:   pkg,
		IsInGoEnv: isInGoEnv,
		GoMod: GoMod{
			Module: module,
			Path:   goMod,
		},
	}, nil
}

func matchModule(data []byte) (string, error) {
	text := string(data)
	re := regexp.MustCompile(`(?m)^\s*module\s+[a-z0-9_/\-.]+$`)
	matches := re.FindAllString(text, -1)
	if len(matches) == 1 {
		target := matches[0]
		index := strings.Index(target, "module")
		return strings.TrimSpace(target[index+6:]), nil
	}

	return "", nil
}
