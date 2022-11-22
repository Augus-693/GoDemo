package main //表示文件所在包是main，在go中，每个文件都必须归属一个包

import "fmt" //表示：引入一个包，包名 fmt, 引入该包后，就可以使用 fmt 包的函数，比如：fmt.Println

/**
 * @Project GoDemo
 * @File 01_hello1.go
 * @Author Augus Lee
 * @Date 2022/9/19 14:28
 * @Version 1.0
 */

func main() { //func 是一个关键字，表示一个函，server 是函数名，是一个主函数，即我们程序的入口。
	fmt.Println("Hello World!") //表示调用 fmt 包的函数 Println 输出 “hello,world”
}
