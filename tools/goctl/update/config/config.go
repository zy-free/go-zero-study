package config

import "go-zero-study/core/logx"

type Config struct {
	logx.LogConf
	ListenOn string
	FileDir  string
	FilePath string
}
