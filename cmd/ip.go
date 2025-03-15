package cmd

import (
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// ipCmd represents the ip command
var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "get ip information",
	Long: `A command line tool for get ip
website: https://ip.sb
IP查询 IPv4.IPv6`,
	Run: func(cmd *cobra.Command, args []string) {
		url := "https://api.ip.sb/ip"
		if ipv4 {
			url = "https://api-ipv4.ip.sb/ip"
		}
		if ipv6 {
			url = "https://api-ipv6.ip.sb/ip"
		}
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Add("User-Agent", "Mozilla")
		resp, err := http.DefaultClient.Do(req)
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
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ipCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	ipCmd.Flags().BoolVarP(&ipv4, "ipv4", "4", false, "get ipv4")
	ipCmd.Flags().BoolVarP(&ipv6, "ipv6", "6", false, "get ipv6")
}
