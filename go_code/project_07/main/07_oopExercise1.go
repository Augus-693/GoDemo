package main

import (
	"fmt"
	"strconv"
)

/**
 * @Project GoDemo
 * @File 07_oopExercise1.go
 * @Author Augus Lee
 * @Date 2022/10/14 14:16
 * @Version 1.0
 */

/*
1) 编写一个 Student 结构体，包含 name、gender、age、id、score 字段，分别为 string、string、int、int、float64 类型。
2) 结构体中声明一个 say 方法，返回 string 类型，方法返回信息中包含所有字段值。
3) 在main 方法中，创建 Student 结构体实例(变量)，并访问 say 方法，并将调用结果打印输出。
*/

/**
 * @StructName Student
 * @Description: 定义一个学生结构体
 **/

type Student struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

/**
 * @functionName say
 * @author Augus Lee
 * @description 实现说这个方法
 * @date 2022-10-18 15:33:29
 * @receive student *Student
 * @return string
 **/
func (student *Student) say() string {
	return student.name + " " + student.gender + " " +
		strconv.FormatInt(int64(student.age), 10) + " " +
		strconv.Itoa(student.id) + " " +
		strconv.FormatFloat(student.score, 'f', 10, 64)
}

/**
 * @functionName server
 * @author Augus Lee
 * @description
 * @date 2022-10-18 15:33:55
 **/
func main() {
	student := Student{name: "Jack", gender: "male", age: 22, id: 101, score: 98.5}
	str := student.say()
	fmt.Println(str)

}
