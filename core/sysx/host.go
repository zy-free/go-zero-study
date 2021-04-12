package sysx

import (
	"os"

	"go-zero-study/core/stringx"
)

var hostname string

func init() {
	var err error
	hostname, err = os.Hostname()
	if err != nil {
		hostname = stringx.RandId()
	}
}

func Hostname() string {
	return hostname
}
