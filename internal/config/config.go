package config

import (
	"errors"
	"github.com/BurntSushi/toml"
	"os"
	"path"
)

var configPath = path.Join("snapshot-manager", "config")

type Config struct {
	SnapRoot string
}

func Get() (Config, error) {
	var conf Config

	raw, err := read()
	if err != nil {
		return conf, err
	}

	_, err = toml.Decode(raw, &conf)
	return conf, err
}

func read() (string, error) {
	etc := path.Join("/etc", configPath)
	bytes, err := os.ReadFile(etc)
	if err == nil {
		return string(bytes), nil
	}

	return "", errors.New("No config found")
}
