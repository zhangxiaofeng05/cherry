package cmd

import (
	"github.com/spf13/cobra"
	"io"
	"log"
	"net/http"
)

// ipCmd represents the ip command
var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "get ip information",
	Long: `A command line tool for get ip
website: https://ipw.cn
IP查询 IPv4.IPv6`,
	Run: func(cmd *cobra.Command, args []string) {
		url := "https://test.ipw.cn/"
		if ipv4 {
			url = "https://4.ipw.cn/"
		}
		if ipv6 {
			url = "https://6.ipw.cn/"
		}
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(string(bytes))
	},
}

var (
	ipv4 bool
	ipv6 bool
)

func init() {
	rootCmd.AddCommand(ipCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ipCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	ipCmd.Flags().BoolVarP(&ipv4, "ipv4", "4", false, "get ipv4")
	ipCmd.Flags().BoolVarP(&ipv6, "ipv6", "6", false, "get ipv6")
}
