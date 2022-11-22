package main

import (
	"fmt"
	"time"
)

/**
 * @Project GoDemo
 * @File    11_primeStatistic.go
 * @Author  Augus Lee
 * @Date    2022/10/21 15:29
 * @GoVersion go1.19.2
 * @Version 1.0
 */

/**
 * @functionName putNum
 * @author Augus Lee
 * @description 向intChan放入数据
 * @date 2022-10-21 15:30:37
 * @param intChan chan int
 **/
func putNum(intChan chan int) {
	for i := 1; i <= 80000; i++ {
		intChan <- i
	}
	close(intChan)
}

/**
 * @functionName primeNum
 * @author Augus Lee
 * @description 从intChan读取数据，若为素数，放到 primeChan
 * @date 2022-10-21 15:32:35
 * @param intChan chan int
 * @param primeChan chan int
 * @param exitChan chan bool
 **/
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool
	for {
		time.Sleep(time.Millisecond * 10)
		num, ok := <-intChan
		if !ok { //intChan取不到
			break
		}
		flag = true //假定是素数
		//判断num 是不是素数
		for i := 2; i <= num; i++ {
			if num%i == 0 { //说明num不是素数
				flag = false
				break
			}
		}
		if flag {
			//将数据放入  primeChan
			primeChan <- num
		}
	}
	fmt.Println("有一个 primeChan 协程因为取不到数据，退出")
	//向 exitChan 写入 true
	exitChan <- true
}

func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 20000) //放入结果
	exitChan := make(chan bool, 4)     //标识退出的管道

	go putNum(intChan) //开启协程放入数据
	//开启四个协程，判断数据是否为素数
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	//主线程
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(primeChan)
	}()
	//遍历 primeChan ，将结果取出
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		//将结果输出
		fmt.Printf("素数 = %v\n", res)
	}
	fmt.Println("main线程退出")
}
