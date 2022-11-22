package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
)

/**
 * @Project GoDemo
 * @File    09_fileWriteEruptSimultaneously.go
 * @Author  Augus Lee
 * @Date    2022/10/17 17:32
 * @GoVersion go1.19.2
 * @Version 1.0
 */

// 产生随机数
// 产生随机数并且将 data 写入到 channel 中，
// 之后通过调用 waitGroup 的 Done 方法来通知任务已经完成
func produce(data chan int, wg *sync.WaitGroup) {
	n := rand.Intn(999)
	data <- n
	wg.Done()
}

// 将数据写到文件
// consume 的函数创建了一个名为 concurrent 的文件。
// 然后从 channel 中读取随机数并且写到文件中。
// 一旦读取完成并且将随机数写入文件后，
// 通过往 done 这个 cahnnel 中写入 true 来通知任务已完成
func consume(data chan int, done chan bool) {
	f, err := os.Create("concurrent")
	if err != nil {
		fmt.Println(err)
		return
	}
	for d := range data {
		_, err = fmt.Fprintln(f, d)
		if err != nil {
			fmt.Println(err)
			f.Close()
			done <- false
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		done <- false
		return
	}
	//将 true 写入 done 这个 channel 中，
	//这个时候 server 函数解除阻塞并且打印 File written successfully
	done <- true
}

func main() {
	data := make(chan int) //创建写入和读取数据的 channel
	//创建 done 这个 channel，此 channel 用于消费者 goroutinue 完成任务之后通知 server 函数
	done := make(chan bool)
	//创建 Waitgroup 的实例 wg，用于等待所有生产随机数的 goroutine 完成任务
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ { //使用 for 循环创建 100 个 goroutines
		wg.Add(1)
		go produce(data, &wg)
	}
	go consume(data, done)
	//调用 waitgroup 的 wait() 方法等待所有的 goroutines 完成随机数的生成。
	//然后关闭 channel。当 channel 关闭时，
	//消费者 consume goroutine 已经将所有的随机数写入文件
	go func() {
		wg.Wait()
		close(data)
	}()
	d := <-done
	if d == true {
		fmt.Println("File written successfully")
	} else {
		fmt.Println("File writing failed")
	}
}
