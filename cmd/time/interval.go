package time

import (
	"github.com/spf13/cobra"
	"log"
	"time"
)

var startTimeStr, endTimeStr string

func init() {
	IntervalTimeCalcCmd.PersistentFlags().StringVarP(&startTimeStr, "start", "s", "", "Start time in DateOnly format")
	IntervalTimeCalcCmd.PersistentFlags().StringVarP(&endTimeStr, "end", "e", "", "End time in DateOnly format")
	IntervalTimeCalcCmd.MarkPersistentFlagRequired("start")
	IntervalTimeCalcCmd.MarkPersistentFlagRequired("end")
}

var IntervalTimeCalcCmd = &cobra.Command{
	Use:   "interval",
	Short: "两个时间的间隔",
	Long:  "两个时间的间隔 例如: 2006-01-02",
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
