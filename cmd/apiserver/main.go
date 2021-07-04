package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/alexispell/gopher-school/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	// (Set config to &configPath, flag, fileDir, description
		// defines flags to ./apiserver 
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path")
}

func main() {
	flag.Parse()


	config := apiserver.NewConfig()
	// Decode toml file (dir to file with config, to where put config) 
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	server := apiserver.New(config)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}