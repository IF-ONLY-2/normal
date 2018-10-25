package main

import (
	"time"
)

func main() {
	var timeLocation *time.Location
	timeLocation, _ = time.LoadLocation("PRC")  //定义时区

	/*将时间字符转为时间,指定时区*/
	flag , _ := time.ParseInLocation("2006-01-02 15:04:05", "2018-10-25 11:37:18", timeLocation)

	/*将当前时间转为固定时区*/
	now := time.Now().In(timeLocation)

	now.Before(flag)
	now.After(flag) //范围判断

	now.Add(2*time.Second) //时间计算
	now.Format("2006-01-02 15:04:05") //格式化


}
