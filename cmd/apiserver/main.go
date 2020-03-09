package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/gtmartem/go-http-rest-api/internal/app/apiserver"
	"log"
)


var (
	configPath string
)


func init() {
	flag.StringVar(
		&configPath,
		"config-path",
		"configs/apiserver.toml",
		"path to config file",
	)
}


func main() {
	// parse configs
	flag.Parse()
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	// create server
	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}

