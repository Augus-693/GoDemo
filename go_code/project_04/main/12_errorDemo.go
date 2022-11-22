package main

import (
	"errors"
	"fmt"
)

/**
 * @Project GoDemo
 * @File 12_errorDemo.go
 * @Author Augus Lee
 * @Date 2022/10/10 14:07
 * @Version 1.0
 */

func test01() {
	//使用defer + recover 来捕获和处理异常
	defer func() {
		err := recover() // recover()内置函数，可以捕获到异常
		if err != nil {  // 说明捕获到错误
			fmt.Println("err=", err)
			//这里进行发生错误时的处理
			fmt.Println("发生错误，请及时处理")
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

// 函数去读取以配置文件init.conf的信息
// 如果文件名传入不正确，我们就返回一个自定义的错误
func readConf(name string) (err error) {
	if name == "config.ini" {
		//读取...
		return nil
	} else {
		//返回一个自定义错误
		return errors.New("读取文件错误..")
	}
}

func test02() {
	err := readConf("config2.ini")
	if err != nil {
		//如果读取文件发送错误，就输出这个错误，并终止程序
		panic(err)
	}
	fmt.Println("test02()继续执行....")
}

func main() {

	//测试
	test01()

	//测试自定义错误的使用
	test02()
	fmt.Println("server()下面的代码...")
}
