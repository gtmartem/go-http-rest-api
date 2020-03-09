package main

import (
	"fmt"
	"github.com/gtmartem/go-http-rest-api/cmd/apiserver"
	"github.com/spf13/cobra"
	"os"
)

func main() {

	rootCmd := &cobra.Command{
		Use:   "apiserver",
		Short: "GO HTTP REST API Server",
	}

	rootCmd.AddCommand(
		apiserver.Cmd(),
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
