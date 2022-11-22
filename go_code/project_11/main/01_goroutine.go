package main

import (
	"fmt"
	"strconv"
	"time"
)

/**
 * @Project GoDemo
 * @File    01_goroutine.go
 * @Author  Augus Lee
 * @Date    2022/10/18 18:59
 * @GoVersion go1.19.2
 * @Version 1.0
 */

func test() {
	for i := 0; i < 100; i++ {
		fmt.Println("test() hello world " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {
	go test() //开启了一个协程
	for i := 0; i < 10; i++ {
		fmt.Println("server() hello golang " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
