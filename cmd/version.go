package cmd

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"leetcode-tool/config"
	"os"
	"runtime"
	"runtime/debug"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get application version",
	Args:  cobra.NoArgs,
	Run: func(_ *cobra.Command, args []string) {
		if config.IsDebugMode() {
			info, _ := debug.ReadBuildInfo()
			fmt.Println(info.String())
		}
		fmt.Printf("version: %s (%s)\n", config.Version, config.GetMode())
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
