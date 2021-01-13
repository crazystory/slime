package main

import (
	"flag"
	"github.com/crazystory/slime/app"
)

var (
	config string
)

func main() {
	flag.StringVar(&config, `config`, `config.yaml`, `config path`)
	flag.Parse()

	app.InitConfig(config)
	app.InitLogger()

	//@todo signal
}
