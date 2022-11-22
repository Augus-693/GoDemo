package main

import (
	"fmt"
	"time"
)

/**
 * @Project GoDemo
 * @File 10_timeDemo.go
 * @Author Augus Lee
 * @Date 2022/10/10 12:04
 * @Version 1.0
 */

func main() {
	//获取当前时间
	now := time.Now()
	fmt.Printf("now = %v, now type = %T\n", now, now)

	//获取其他的日期信息
	fmt.Printf("年 = %v\n", now.Year())
	fmt.Printf("月 = %v\n", now.Month)
	fmt.Printf("月 = %v\n", int(now.Month()))
	fmt.Printf("日 = %v\n", now.Day())
	fmt.Printf("时 = %v\n", now.Hour())
	fmt.Printf("分 = %v\n", now.Minute())
	fmt.Printf("秒  = %v\n", now.Second())

	//格式化日期时间
	fmt.Printf("当前年月日 ：%d-%d-%d %d:%d:%d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	dateStr := fmt.Sprintf("当前年月日 ：%d-%d-%d %d:%d:%d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(dateStr)
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("15:04:05"))

	//结合 Sleep 使用时间变量
	//每隔一秒打印一个数字，打印到10退出
	i := 0
	for {
		i++
		fmt.Println(i)
		//休眠
		time.Sleep(time.Second)
		if i == 10 {
			break
		}
	}

	//time 的 Unix 和 UnixNano 的方法
	fmt.Printf("unix时间戳 = %v, unixnono时间戳 = %v\n", now.Unix(), now.UnixNano())
}
