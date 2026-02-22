package model

import (
	"fmt"
	"time"
)

// 传入时间戳，返回时间格式字符串
func UnixToTime(timestamp int) string {
	fmt.Println(timestamp)
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}
