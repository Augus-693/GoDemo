package main

import (
	"flag"
	"fmt"
)

/**
 * @Project GoDemo
 * @File    12_argsDemo.go
 * @Author  Augus Lee
 * @Date    2022/10/18 14:28
 * @GoVersion go1.19.2
 * @Version 1.0
 */

/**
 * @functionName server
 * @author Augus Lee
 * @description 获取命令行参数
 * @date 2022-10-18 15:41:04
 **/
func main() {
	//定义几个变量，用于接受命令行的参数值
	var user string
	var pwd string
	var host string
	var port int

	flag.StringVar(&user, "user", "", "用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码，默认为空")
	flag.StringVar(&host, "host", "localhost", "主机名，默认localhost")
	flag.IntVar(&port, "port", 3306, "端口号，默认3306")
	flag.Parse()
	fmt.Printf("user = %v pwd = %v host = %v port = %v\n", user, pwd, host, port)
}
