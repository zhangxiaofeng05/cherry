package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zhangxiaofeng05/cherry/internal/timer"
	"log"
	"time"
)

var startTimeStr, endTimeStr string

func init() {
	intervalTimeCalcCmd.PersistentFlags().StringVarP(&startTimeStr, "start", "s", "", "Start time in DateOnly format")
	intervalTimeCalcCmd.PersistentFlags().StringVarP(&endTimeStr, "end", "e", "", "End time in DateOnly format")
	intervalTimeCalcCmd.MarkPersistentFlagRequired("start")
	intervalTimeCalcCmd.MarkPersistentFlagRequired("end")
}

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

var intervalTimeCalcCmd = &cobra.Command{
	Use:   "interval",
	Short: "两个时间的间隔",
	Long:  "两个时间的间隔",
	Run: func(cmd *cobra.Command, args []string) {
		startTime, err := time.Parse(time.DateOnly, startTimeStr)
		if err != nil {
			log.Fatalf("Error parsing start time err:%v", err)
		}

		endTime, err := time.Parse(time.DateOnly, endTimeStr)
		if err != nil {
			log.Fatalf("Error parsing end time err: %v", err)
		}

		interval := endTime.Sub(startTime)
		h := interval.Hours() / 24

		log.Printf("Time Interval: %v day", h)
	},
}

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(intervalTimeCalcCmd)
}
