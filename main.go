package main

import (
	"log"

	"github.com/zhangxiaofeng05/cherry/cmd"
)

func main() {
	log.SetFlags(0) // 取消log任何前缀

	cmd.Execute()
}
