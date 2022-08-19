package cmd

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"leetcode-tool/config"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "get version",
	Args:  cobra.NoArgs,
	Run: func(_ *cobra.Command, args []string) {
		fmt.Printf("version: %s\n", config.Version)
		fmt.Printf("build: %s %s/%s\n", runtime.Version(), runtime.GOOS, runtime.GOARCH)

		appPath, err := os.Executable()
		if err != nil {
			fmt.Println("get path:", err)
			return
		}
		fmt.Printf("app path: %s\n", appPath)

		data, err := os.ReadFile(appPath)
		if err != nil {
			fmt.Println("read file:", err)
			return
		}
		hash := sha1.Sum(data)
		fmt.Printf("checksum: %s\n", hex.EncodeToString(hash[:]))
	},
}
