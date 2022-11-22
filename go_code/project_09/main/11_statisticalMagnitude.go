package main

import (
	"fmt"
	"os"
)

/**
 * @Project GoDemo
 * @File    11_statisticalMagnitude.go
 * @Author  Augus Lee
 * @Date    2022/10/18 12:54
 * @GoVersion go1.19.2
 * @Version 1.0
 */

// 定义一个结构体，用于保存统计结果
type CharCount struct {
	ChCount    int // 记录英文个数
	NumCount   int // 记录数字的个数
	SpaceCount int // 记录空格的个数
	OtherCount int // 记录其它字符的个数
}

func main() {
	//定义个CharCount 实例
	var count CharCount

	//开始循环的读取fileName的内容
	var content []byte
	var err error
	for {
		content, err = os.ReadFile("../file/string.txt")
		if err == nil {
			break
		}
	}

	//遍历 str ，进行统计
	for _, v := range content {
		switch {
		case v >= 'a' && v <= 'z':
			fallthrough //穿透
		case v >= 'A' && v <= 'Z':
			count.ChCount++
		case v == ' ' || v == '\t':
			count.SpaceCount++
		case v >= '0' && v <= '9':
			count.NumCount++
		default:
			count.OtherCount++
		}
	}

	//输出统计的结果看看是否正确
	fmt.Printf("字符的个数为=%v 数字的个数为=%v 空格的个数为=%v 其它字符个数=%v",
		count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)
}
