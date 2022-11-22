package main

import (
	"fmt"
	"github.com/GoDemo/go_code/project_07/main/model"
)

/**
 * @Project GoDemo
 * @File 08_factory.go
 * @Author Augus Lee
 * @Date 2022/10/14 15:12
 * @Version 1.0
 */

func main() {
	//定student结构体是首字母小写，我们可以通过工厂模式来解决
	var stu = model.NewStudent("tom", 22, 98.8)
	fmt.Println(*stu)
	fmt.Println("name=", stu.Name, " Age=", stu.Age, " score=", stu.GetScore())
}
