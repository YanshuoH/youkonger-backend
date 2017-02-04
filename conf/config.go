package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/YanshuoH/youkonger/utils"
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
	file, err := utils.GetAbsFilePath(file)
	if err != nil {
		return nil, err
	}

	c := &config{}
	if _, err := toml.DecodeFile(file, c); err != nil {
		return nil, err
	}

	Config = c
	return Config, nil
}
