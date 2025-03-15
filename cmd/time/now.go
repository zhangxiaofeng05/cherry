package time

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/zhangxiaofeng05/cherry/internal/timer"
)

var NowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "获取当前时间",
	Long:  "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		log.Printf("当前时间: %s \t unix(10,秒): %d \t unixMilli(13,毫秒): %d", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix(), nowTime.UnixMilli())
	},
}
