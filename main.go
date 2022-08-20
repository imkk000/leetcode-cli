package main

import (
	"fmt"
	"leetcode-tool/cmd"
	"leetcode-tool/config"

	"github.com/spf13/cobra"
)

func main() {
	cd := cobra.Command{}

	// implementing command
	if config.IsDebugMode() {
		cd.AddCommand(
			cmd.GenerateFileCmd,
			cmd.TestCmd,
		)
	}

	// publish command
	cd.AddCommand(
		cmd.VersionCmd,
		cmd.RunCmd,
		cmd.InfoCmd,
		cmd.ParseFileCmd,
		cmd.TestNetCmd,
	)

	if err := cd.Execute(); err != nil {
		fmt.Println("run:", err)
	}
}
