package config

import "github.com/hashicorp/go-hclog"

var Logger = hclog.New(&hclog.LoggerOptions{
	Name:  "api",
	Level: hclog.LevelFromString("DEBUG"),
})

