package main

import (
	"fmt"
	"leetcode-tool/cmd"

	"github.com/spf13/cobra"
)

func main() {
	cd := cobra.Command{}
	cd.AddCommand(
		cmd.VersionCmd,
		cmd.RunCmd,
		cmd.InfoCmd,
		cmd.GenerateFileCmd,
		cmd.TestNetCmd,
		cmd.TestCmd,
	)
	if err := cd.Execute(); err != nil {
		fmt.Println("run:", err)
	}
}
