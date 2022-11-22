package main

import "fmt"

/**
 * @Project GoDemo
 * @File    06_channelIntroduction.go
 * @Author  Augus Lee
 * @Date    2022/10/21 12:56
 * @GoVersion go1.19.2
 * @Version 1.0
 */

func main() {
	//演示管道的使用
	//创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int, 3)

	//intChan的信息
	fmt.Printf("intChan的值 = %v, 本身的地址 = %p\n", intChan, &intChan)

	//向管道写入数据
	intChan <- 10
	num := 211
	intChan <- num
	intChan <- 50

	//管道的长度和cap(容量)
	fmt.Printf("intChan len = %v, cap = %v\n", len(intChan), cap(intChan))

	//从管道读取数据
	var num2 int
	num2 = <-intChan
	fmt.Println("num2 =", num2)
	fmt.Printf("channel len = %v, cap = %v\n", len(intChan), cap(intChan))

	//在没有使用协程的情况下，如果管道内数据以全部取出，再取就会报 deadlock
	num3 := <-intChan
	num4 := <-intChan
	num5 := <-intChan
	fmt.Printf("num3 = %v, num4 = %v, num5 = %v\n", num3, num4, num5)

}
