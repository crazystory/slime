package app

import (
	"fmt"
	"github.com/crazystory/slime/modules/log"
	"github.com/spf13/viper"
)

func InitConfig(config string) {
	viper.SetConfigFile(config)

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf(`error while loading config file [%s]: %s`, config, err.Error()))
	}
}

func InitLogger() {
	if err := log.Init(viper.GetString(`log.default`)); err != nil {
		panic(fmt.Sprintf(`log error: %s`, err.Error()))
	}
}
