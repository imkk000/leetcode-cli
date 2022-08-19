package cmd

import (
	"fmt"
	"leetcode-tool/client"
	"leetcode-tool/config"
	"net/http/httputil"
	"net/url"

	"github.com/spf13/cobra"
)

var TestNetCmd = &cobra.Command{
	Use:   "testnet",
	Short: "Test send HEAD to server",
	Run: func(_ *cobra.Command, _ []string) {
		configDir, err := config.GetConfigDir()
		if err != nil {
			fmt.Println("get config dir:", err)
			return
		}
		fmt.Printf("cookies path: %s\n", configDir)
		dumpValues(config.ReadCookies(configDir))
		fmt.Println()

		fmt.Printf("client: %+v\n\n", client.DefaultClient)

		resp, err := client.Head("")
		if err != nil {
			fmt.Println("http get:", err)
			return
		}

		dumpData, err := httputil.DumpRequest(resp.Request, true)
		if err != nil {
			fmt.Println("dump req:", err)
			return
		}
		fmt.Println("request:")
		fmt.Println(string(dumpData))

		dumpData, err = httputil.DumpResponse(resp, true)
		if err != nil {
			fmt.Println("dump resp:", err)
			return
		}
		fmt.Println("response:")
		fmt.Println(string(dumpData))
	},
}

func dumpValues(v url.Values) {
	for k := range v {
		fmt.Printf("- %s= %s\n", k, v.Get(k))
	}
}
