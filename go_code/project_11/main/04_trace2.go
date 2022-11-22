package main

import (
	"fmt"
	"time"
)

/**
 * @Project GoDemo
 * @File    04_trace2.go
 * @Author  Augus Lee
 * @Date    2022/10/19 11:04
 * @GoVersion go1.19.2
 * @Version 1.0
 */

func main() {
	for i := 0; i < 55; i++ {
		time.Sleep(time.Second)
		fmt.Println("Hello World!")
	}
}
