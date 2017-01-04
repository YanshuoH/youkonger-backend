package conf

import (
	"gopkg.in/gcfg.v1"
	"path/filepath"
)

var Config *config

type config struct {
	AppConf appConf
}

type appConf struct {
	GinMode string
}

func Setup(file string) *config {
	if !filepath.IsAbs(file) {
		f, e := filepath.Abs(file)
		if e != nil {
			panic(e)
		}
		file = f
	}

	c := &config{}
	if err := gcfg.ReadFileInto(c, file); err != nil {
		panic(err)
	}

	Config = c
	return Config
}
