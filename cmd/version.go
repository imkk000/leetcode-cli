package cmd

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "get version",
	Args:  cobra.NoArgs,
	Run: func(_ *cobra.Command, args []string) {
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
		hash := sha256.Sum256(data)
		fmt.Printf("checksum: %s\n", hex.EncodeToString(hash[:]))
	},
}
