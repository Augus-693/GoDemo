package main

import "fmt"

/**
 * @Project GoDemo
 * @File 07_oopExercise2.go
 * @Author Augus Lee
 * @Date 2022/10/14 14:33
 * @Version 1.0
 */

/*
1) 编程创建一个 Box 结构体，在其中声明三个字段表示一个立方体的长、宽和高，长宽高要从终端获取
2) 声明一个方法获取立方体的体积。
3) 创建一个 Box 结构体变量，打印给定尺寸的立方体的体积
*/

type Box struct {
	length float64
	width  float64
	height float64
}

/**
 * @functionName getVolume
 * @author Augus Lee
 * @description
 * @date 2022-10-18 15:31:58
 * @receive box Box
 * @return float64
 **/
func (box Box) getVolume() float64 {
	res := box.length * box.width * box.height
	return res
}

/**
 * @functionName server
 * @author Augus Lee
 * @description
 * @date 2022-10-18 15:32:41
 **/
func main() {
	box := new(Box)
	fmt.Printf("请输入Box的长宽高:")
	fmt.Scanf("%v %v %v\n", &box.length, &box.width, &box.height)
	volume := box.getVolume()
	fmt.Println(volume)

}
