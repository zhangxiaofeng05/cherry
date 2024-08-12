package cmd

import (
	"fmt"
	"log"
	"net/url"
	"unicode"

	"github.com/spf13/cobra"
)

var urlCmd = &cobra.Command{
	Use:   "url",
	Short: "url encode and decode",
	Long:  `url encode non ascii chars and decode`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if encode {
			res := EncodeNonASCIIChars(args[0])
			log.Printf("encode result: %s", res)
		} else {
			res, err := DecodeEncodedURL(args[0])
			if err != nil {
				log.Fatalf("decode error: %v", err)
			}
			log.Printf("decode result: %s", res)
		}
	},
}

// EncodeNonASCIIChars 对字符串中的非 ASCII 字符进行百分比编码
func EncodeNonASCIIChars(input string) string {
	var encodedPath string
	for _, r := range input {
		if r > 127 || !unicode.IsPrint(r) { // 非 ASCII 或不可打印字符
			encodedPath += url.QueryEscape(string(r))
		} else {
			encodedPath += string(r)
		}
	}
	return encodedPath
}

// DecodeEncodedURL 解码已编码的 URL
func DecodeEncodedURL(encodedURL string) (string, error) {
	decodedPath, err := url.QueryUnescape(encodedURL)
	if err != nil {
		return "", fmt.Errorf("解码路径失败: %v", err)
	}
	return decodedPath, nil
}

var (
	encode bool
)

func init() {
	urlCmd.Flags().BoolVarP(&encode, "encode", "e", false, "encode")
}
