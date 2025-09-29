package main

import "flag"

var (
	configPath string
)

func Init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
	flag.Parse()
}
