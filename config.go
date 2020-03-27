package main

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Debug    bool
	Secret   string
	Botid    int
	Bottoken string
	Userid   string
}

func loadConfig(path string) (Config, error) {
	var conf Config

	_, err := toml.DecodeFile(path, &conf)
	if err != nil {
		return conf, err
	}

	return conf, nil
}
