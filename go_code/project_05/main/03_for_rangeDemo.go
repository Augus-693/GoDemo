package main

import "fmt"

/**
 * @Project GoDemo
 * @File 03_for_rangeDemo.go
 * @Author Augus Lee
 * @Date 2022/10/10 18:53
 * @Version 1.0
 */

func main() {
	//演示 for-range 遍历数组
	heroes := [...]string{"宋江", "吴用", "卢俊义"}
	for i, v := range heroes {
		fmt.Printf("i = %v, v = %v\n", i, v)
		fmt.Printf("heroes[%d] = %v\n", i, heroes[i])
	}

	for _, v := range heroes {
		fmt.Printf("元素的值 = %v\n", v)
	}
}
