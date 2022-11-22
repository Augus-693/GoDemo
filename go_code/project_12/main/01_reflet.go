package main

import (
	"fmt"
	"reflect"
)

/**
 * @Project GoDemo
 * @File    01_reflet.go
 * @Author  Augus Lee
 * @Date    2022/10/26 15:36
 * @GoVersion go1.19.2
 * @Version 1.0
 */

/**
 * @functionName reflectTest01
 * @author Augus Lee
 * @description 演示对普通数据类型反射
 * @date 2022-10-26 15:38:37
 * @param b interface{}
 **/
func reflectTest01(b interface{}) {
	//通过反射获取传入变量的 type、kind、值
	//1.先获取到 reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rType=", rType)

	//2.获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	n2 := 2 + rVal.Int()
	fmt.Println("n2=", n2)
	fmt.Printf("rVal=%v,rVal type=%T\n", rVal, rVal)

	//将rVal 转成 interface{}
	iV := rVal.Interface()
	//将interface{}通过断言转成需要的类型
	num2 := iV.(int)
	fmt.Println("num2=", num2)
}

/**
 * @functionName reflectTest02
 * @author Augus Lee
 * @description 演示对结构体的反射
 * @date 2022-10-26 15:52:53
 * @param b interface{}
 **/
func reflectTest02(b interface{}) {

	//通过反射获取的传入的变量的 type , kind, 值
	//1. 先获取到 reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rType=", rType)

	//2. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)

	//3. 获取 变量对应的Kind
	//(1) rVal.Kind() ==>
	kind1 := rVal.Kind()
	//(2) rTyp.Kind() ==>
	kind2 := rType.Kind()
	fmt.Printf("kind =%v kind=%v\n", kind1, kind2)

	//下面我们将 rVal 转成 interface{}
	iV := rVal.Interface()
	fmt.Printf("iv=%v iv type=%T \n", iV, iV)
	//将 interface{} 通过断言转成需要的类型
	//这里，我们就简单使用了一带检测的类型断言.
	//同学们可以使用 swtich 的断言形式来做的更加的灵活
	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("stu.Name=%v\n", stu.Name)
	}

}

type Student struct {
	Name string
	Age  int
}

func main() {
	//演示对（基本数据类型、interface{}、reflect.Value）进行反射的基本操作
	//先定义一个int
	var num int = 100
	reflectTest01(num)

	//定义一个Student的实例
	stu := Student{
		Name: "tom",
		Age:  20,
	}
	reflectTest02(stu)
}
