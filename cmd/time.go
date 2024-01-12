package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zhangxiaofeng05/cherry/cmd/time"
)

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式处理",
	Long:  "时间格式处理",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	timeCmd.AddCommand(time.NowTimeCmd)
	timeCmd.AddCommand(time.IntervalTimeCalcCmd)
}
