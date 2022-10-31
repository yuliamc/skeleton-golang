package main

import (
	"os"

	"github.com/spf13/cobra"

	"modalrakyat/skeleton-golang/cmd/apiserver/app"
	"modalrakyat/skeleton-golang/config"
)

var (
	rootCMD = &cobra.Command{
		Short: "skeleton-golang",
	}

	configCMD = &cobra.Command{
		Use:   "config",
		Short: "Show settings",
		Run: func(*cobra.Command, []string) {
			config.Show()
		},
	}

	serverCMD = &cobra.Command{
		Use:   "server",
		Short: "Run application server",
		Run: func(*cobra.Command, []string) {
			app.Run()
		},
	}
)

func main() {
	cobra.OnInitialize(config.Init)

	// Regist
	rootCMD.AddCommand(configCMD)
	rootCMD.AddCommand(serverCMD)
	if err := rootCMD.Execute(); err != nil {
		os.Exit(1)
	}
}
