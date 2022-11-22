package main

import "fmt"

/**
 * @Project GoDemo
 * @File 02_sequentialFind.go
 * @Author Augus Lee
 * @Date 2022/10/13 14:08
 * @Version 1.0
 */

func main() {
	names := [4]string{"golang", "java", "python", "php"}
	var name = ""
	fmt.Printf("请输入要查找的语言：")
	fmt.Scanln(&name)

	//顺序查找：第一种方式
	for i := 0; i < len(names); i++ {
		if name == names[i] {
			fmt.Printf("找到%v,下标%v\n", name, i)
			break
		} else if i == len(names)-1 {
			fmt.Printf("没有找到%v\n", name)
		}
	}

	//顺序查找：第二种方法
	index := -1
	for i := 0; i < len(names); i++ {
		if name == names[i] {
			index = i
			break
		}
	}
	if index != -1 {
		fmt.Printf("找到%v,下标%v\n", name, index)
	} else {
		fmt.Println("没有找到", name)
	}
}
