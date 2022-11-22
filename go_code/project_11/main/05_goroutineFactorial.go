package main

import (
	"fmt"
	"sync"
	"time"
)

/**
 * @Project GoDemo
 * @File    05_goroutineFactorial.go
 * @Author  Augus Lee
 * @Date    2022/10/19 13:19
 * @GoVersion go1.19.2
 * @Version 1.0
 */

var (
	myMap = make(map[int]int, 10)
	//声明一个全局的互斥锁
	lock sync.Mutex
)

/**
 * @functionName factorial
 * @author Augus Lee
 * @description 计算n！,将结果放入map
 * @date 2022-10-19 13:58:52
 * @param n int
 **/
func factorial(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	lock.Lock() //加锁
	myMap[n] = res
	lock.Unlock() //解锁
}

func main() {
	for i := 1; i <= 20; i++ {
		go factorial(i)
	}
	time.Sleep(time.Second * 5)
	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("myMap[%d]: %v\n", i, v)
		//fmt.Printf("%v! = %v\n", i+1, v)
	}
	lock.Unlock()
}
