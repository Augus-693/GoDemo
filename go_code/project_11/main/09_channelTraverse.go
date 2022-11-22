package main

import "fmt"

/**
 * @Project GoDemo
 * @File    09_channelTraverse.go
 * @Author  Augus Lee
 * @Date    2022/10/21 14:43
 * @GoVersion go1.19.2
 * @Version 1.0
 */

func main() {
	var intChan = make(chan int, 100)
	for i := 0; i < 100; i++ { //往管道写入100个数据
		intChan <- i * 2
	}
	//必须要关闭管道，否则报错deadlock
	close(intChan)
	//不能使用fori常规遍历
	for v := range intChan {
		fmt.Printf("v = %v ", v)
	}
}
