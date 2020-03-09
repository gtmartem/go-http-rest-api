package apiserver

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/gtmartem/go-http-rest-api/internal/app/apiserver"
	"github.com/spf13/cobra"
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


func startCmd(cmd *cobra.Command, args []string) {
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


func Cmd() *cobra.Command {
	return &cobra.Command{
		Use:   "apiserver",
		Short: "Use this command to run go http rest api server",
		Run:   startCmd,
	}
}
