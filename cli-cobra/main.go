package main

import (
	"cli_cobra/cmd"
	"cli_cobra/cmd/keys"
	"cli_cobra/cmd/signatures"
	"cli_cobra/logger"
)

func main() {
	rootCmd := cmd.RootCmd()
	keys.Init(rootCmd)
	signatures.Init(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		logger.HaltOnErr(err, "failed to execute command")
	}
}
