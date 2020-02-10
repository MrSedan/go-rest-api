package main

import (
	"flag"
	"log"

	"github.com/MrSedan/go-rest-api/internal/app/apiserver"
	"github.com/pelletier/go-toml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	data, err := toml.LoadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}
	data.Unmarshal(config)

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}

}
