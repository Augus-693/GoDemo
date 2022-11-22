package main

import (
	"fmt"
	"github.com/GoDemo/go_code/project_08/model"
)

/**
 * @Project GoDemo
 * @File 02_encapsulation.go
 * @Author Augus Lee
 * @Date 2022/10/14 18:17
 * @Version 1.0
 */

func main() {
	p := model.NewPerson("John")
	p.SetAge(18)
	p.SetSal(15000)
	fmt.Println(p)
	fmt.Println(p.Name, " age = ", p.GetAge(), " sal = ", p.GetSal())
}
