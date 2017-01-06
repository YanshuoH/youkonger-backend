package conf

import (
	"github.com/BurntSushi/toml"
	"path/filepath"
)

var Config *config

type config struct {
	AppConf appConf
	DbConf  dbConf
}

type appConf struct {
	GinMode string
}

type dbConf struct {
	Dsn string
}

func Setup(file string) (*config, error) {
	if !filepath.IsAbs(file) {
		f, e := filepath.Abs(file)
		if e != nil {
			return nil, e
		}
		file = f
	}

	c := &config{}
	if _, err := toml.DecodeFile(file, c); err != nil {
		return nil, err
	}

	Config = c
	return Config, nil
}
