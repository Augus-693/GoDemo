package main

import "fmt"

/**
 * @Project GoDemo
 * @File    08_channelClose.go
 * @Author  Augus Lee
 * @Date    2022/10/21 14:40
 * @GoVersion go1.19.2
 * @Version 1.0
 */

func main() {
	var intChan = make(chan int, 3)
	intChan <- 10
	intChan <- 20
	intChan <- 30
	close(intChan) //关闭管道，此时不能写入数据，但依旧可以读取
	//因为 intChan 的容量是3，再存放会报 deadlock
	num1 := <-intChan
	num2 := <-intChan
	num3 := <-intChan
	//因为这时 intChan 已经没有数据，再取会报 deadlock
	fmt.Printf("num1: %v, num2: %v, num3: %v\n", num1, num2, num3)
}
