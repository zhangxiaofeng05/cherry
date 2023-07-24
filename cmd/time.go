package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zhangxiaofeng05/cherry/internal/timer"
	"log"
)

var (
	calculateTime string
	duration      string
)

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式处理",
	Long:  "时间格式处理",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(timeCmd)
}

var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "获取当前时间",
	Long:  "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		log.Printf("输出结果: %s, %d", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix())
	},
}

func init() {
	timeCmd.AddCommand(nowTimeCmd)
}
