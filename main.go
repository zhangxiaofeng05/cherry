package main

import (
	"github.com/zhangxiaofeng05/cherry/cmd"
	"log"
)

func main() {
	log.SetFlags(0) // 取消log任何前缀

	cmd.Execute()
}
