package model

import (
	"time"
)

// TimeUtil 时间工具结构体
type TimeUtil struct{}

// TimestampToDate 将时间戳转换为日期字符串
// timestamp: 秒级或毫秒级时间戳
// layout: 日期格式，如 "2006-01-02 15:04:05"
func (t TimeUtil) TimestampToDate(timestamp int64, layout string) string {
	// 判断是秒级还是毫秒级时间戳
	if timestamp > 1e10 {
		timestamp = timestamp / 1000
	}
	return time.Unix(timestamp, 0).Format(layout)
}

// TimestampToDateDefault 使用默认格式转换时间戳
// 默认格式: 2006-01-02 15:04:05
func (t TimeUtil) TimestampToDateDefault(timestamp int64) string {
	return t.TimestampToDate(timestamp, "2006-01-02 15:04:05")
}

// TimestampToDateOnly 只返回日期部分
// 格式: 2006-01-02
func (t TimeUtil) TimestampToDateOnly(timestamp int64) string {
	return t.TimestampToDate(timestamp, "2006-01-02")
}

// NowTimestamp 获取当前时间戳（秒级）
func (t TimeUtil) NowTimestamp() int64 {
	return time.Now().Unix()
}

// NowTimestampMilli 获取当前时间戳（毫秒级）
func (t TimeUtil) NowTimestampMilli() int64 {
	return time.Now().UnixMilli()
}
