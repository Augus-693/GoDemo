package main

import "fmt"

/**
 * @Project GoDemo
 * @File 09_string.go
 * @Author Augus Lee
 * @Date 2022/9/20 19:05
 * @Version 1.0
 */

func main() {
	var address string = "河南省南阳市宛城区南阳理工学院"
	fmt.Println(address)

	//双引号, 会识别转义字符
	str := "abc\ndef"
	fmt.Println(str)

	//以字符串的原生形式输出，包括换行和特殊字符，可以实现防止攻击、输出源代码等效果
	code := `
var address string = "河南省南阳市宛城区南阳理工学院"
fmt.Println(address)

str := "abc\ndef"
fmt.Println(str)
`
	fmt.Println(code)

	//字符串拼接
	str1 := "hello" + " and "
	str1 += "world"
	fmt.Println(str1)

	//当一行字符串太长时，需要使用到多行字符串，可以分行写，注意需要将 + 号保留在上一行
	str3 := "hello" + "world" + "hello" + "world" + "hello" + "world" +
		"hello" + "world" + "hello" + "world" + "hello" + "world" + "hello" +
		"world" + "hello" + "world" + "hello" + "world" + "hello" + "world" +
		"hello" + "world" + "hello" + "world" + "hello" + "world"
	fmt.Println(str3)

}
