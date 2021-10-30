package main

import (
	"fmt"
	"time"
)

/**
时间和日期
*/

var week time.Duration

func main() {

	now := time.Now()
	fmt.Printf("当前时间为：%v \n", now)
	day := now.Day()
	fmt.Printf("日为：%v \n", day)
	minute := now.Minute()
	fmt.Printf("分为：%v \n", minute)

	//格式化时间
	fmt.Printf("%02d.%02d.%4d\n", now.Day(), now.Month(), now.Year())
	fmt.Printf("%4d-%02d-%02d\n", now.Year(), now.Month(), now.Day())
	fmt.Println(now.Format("2006-01-02 15:01:05"))

	week = 60 * 60 * 24 * 7 * 1e9 // must be in nanosec
	week_from_now := now.Add(week)
	fmt.Println(week_from_now) // Wed Dec 28 08:52:14 +0000 UTC 2011
}
