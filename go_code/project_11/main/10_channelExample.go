package main

import "fmt"

/**
 * @Project GoDemo
 * @File    10_channelExample.go
 * @Author  Augus Lee
 * @Date    2022/10/21 14:48
 * @GoVersion go1.19.2
 * @Version 1.0
 */

/**
 * @functionName writeData
 * @author Augus Lee
 * @description 写入数据
 * @date 2022-10-21 15:23:15
 * @param intChan chan int
 **/
func writeData(intChan chan int) {
	for i := 0; i < 50; i++ {
		intChan <- i
		fmt.Println("writeData: ", i)
	}
	close(intChan)
}

/**
 * @functionName readData
 * @author Augus Lee
 * @description 读取数据
 * @date 2022-10-21 15:23:32
 * @param intChan chan int
 * @param exitChan chan bool
 **/
func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("readData读到数据 = %v\n", v)
	}
	exitChan <- true
	close(exitChan)
}

func main() {
	//创建两个管道
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}
