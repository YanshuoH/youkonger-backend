package conf

import (
	"gopkg.in/gcfg.v1"
	"path/filepath"
)

var Config *config

type config struct {
	AppConf appConf
	DbConf dbConf
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
	if err := gcfg.ReadFileInto(c, file); err != nil {
		return nil, err
	}

	Config = c
	return Config, nil
}
