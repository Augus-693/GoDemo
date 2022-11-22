package main

import (
	"fmt"
)

/**
 * @Project GoDemo
 * @File    07_channelCase.go
 * @Author  Augus Lee
 * @Date    2022/10/21 13:55
 * @GoVersion go1.19.2
 * @Version 1.0
 */

type Cat struct {
	Name string
	Age  int
}

/**
 * @functionName intChan
 * @author Augus Lee
 * @description 创建一个intChan，最多可以存放三个int数据，演示存放三个数据到intChan，然后再取出这三个int数据
 * @date 2022-10-21 14:05:54
 **/
func intChan() {
	var intChan = make(chan int, 3)
	intChan <- 10
	intChan <- 20
	intChan <- 30
	//因为 intChan 的容量是3，再存放会报 deadlock
	num1 := <-intChan
	num2 := <-intChan
	num3 := <-intChan
	//因为这时 intChan 已经没有数据，再取会报 deadlock
	fmt.Printf("num1: %v, num2: %v, num3: %v\n", num1, num2, num3)
}

/**
 * @functionName mapChan
 * @author Augus Lee
 * @description 创建一个mapChan，最多可以存放10个map[string]string的key-value，演示写入和读取
 * @date 2022-10-21 14:12:59
 **/
func mapChan() {
	var mapChan = make(chan map[string]string, 10)
	m1 := make(map[string]string, 20)
	m1["city1"] = "北京"
	m1["city2"] = "上海"

	m2 := make(map[string]string, 20)
	m2["hero1"] = "宋江"
	m2["hero2"] = "吴用"
	//。。。
	mapChan <- m1
	mapChan <- m2

	m3 := make(map[string]string, 20)
	m3 = <-mapChan
	m4 := make(map[string]string, 20)
	m4 = <-mapChan
	fmt.Printf("m3 = %v, m4 = %v\n", m3, m4)
}

/**
 * @functionName catChan
 * @author Augus Lee
 * @description 创建一个catChan，最多可以存放10个Cat结构体，演示写入和读取
 * @date 2022-10-21 14:19:11
 **/
func catChan() {
	var catChan = make(chan Cat, 10)
	cat1 := Cat{Name: "Tom", Age: 11}
	cat2 := Cat{Name: "Jack", Age: 12}
	catChan <- cat1
	catChan <- cat2

	cat3 := <-catChan
	cat4 := <-catChan
	fmt.Printf("cat3 = %v, cat4 = %v\n", cat3, cat4)
}

/**
 * @functionName catChan2
 * @author Augus Lee
 * @description 创建一个catChan，最多可以存放10个*Cat结构体，演示写入和读取
 * @date 2022-10-21 14:20:15
 **/
func catChan2() {
	var catChan = make(chan *Cat, 10)
	cat1 := Cat{Name: "Tom", Age: 11}
	cat2 := Cat{Name: "Jack", Age: 12}
	catChan <- &cat1
	catChan <- &cat2

	cat3 := <-catChan
	cat4 := <-catChan
	fmt.Printf("cat3 = %v, cat4 = %v\n", cat3, cat4)
}

/**
 * @functionName allChan
 * @author Augus Lee
 * @description 创建一个allChan，最多可以存放10个任意数据类型变量，演示写入和读取
 * @date 2022-10-21 14:24:21
 **/
func allChan() {
	var allChan = make(chan interface{}, 10)
	cat1 := Cat{Name: "Tom", Age: 11}
	cat2 := Cat{Name: "Jack", Age: 12}
	allChan <- cat1
	allChan <- cat2
	allChan <- 10
	allChan <- "jack"

	cat3 := <-allChan
	cat4 := <-allChan
	v1 := <-allChan
	v2 := <-allChan
	fmt.Printf("cat3: %v, cat4: %v, v1: %v, v2: %v\n", cat3, cat4, v1, v2)
}

func main() {
	intChan()
	mapChan()
	catChan()
	catChan2()
	allChan()
}
