package main

import "fmt"

/**
 * @Project GoDemo
 * @File 07_sliceOfMap.go
 * @Author Augus Lee
 * @Date 2022/10/13 20:08
 * @Version 1.0
 */

func main() {
	//声明一个Map切片
	var students []map[string]string
	students = make([]map[string]string, 2)

	//添加学生信息
	if students[0] == nil {
		students[0] = make(map[string]string, 2)
		students[0]["name"] = "张三"
		students[0]["age"] = "21"
	}
	if students[1] == nil {
		students[1] = make(map[string]string, 2)
		students[1]["name"] = "李四"
		students[1]["age"] = "20"
	}
	//使用切片的append函数，可以动态增加student
	//定义一个student 信息
	newStudent := map[string]string{
		"name": "王五",
		"age":  "22",
	}
	students = append(students, newStudent)
	fmt.Println(students)
}
