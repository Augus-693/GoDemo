package main

import "fmt"

/**
 * @Project GoDemo
 * @File 05_mapStatement.go
 * @Author Augus Lee
 * @Date 2022/10/13 15:29
 * @Version 1.0
 */

func main() {
	//map的声明/使用：第一种方式
	var a map[string]string
	//在使用map前，需要先make，make的作用是给map分配数据空间
	a = make(map[string]string, 10)
	a["no1"] = "golang"
	a["no2"] = "java"
	a["no3"] = "python"
	a["no4"] = "php"
	fmt.Println(a)

	//第二种方式
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "上海"
	cities["no3"] = "深圳"
	cities["no4"] = "天津"
	fmt.Println(cities)

	//第三种方式
	heros := map[string]string{
		"hero1": "宋江",
		"hero2": "吴用",
		"hero3": "卢俊义",
	}
	heros["hero4"] = "林冲"
	fmt.Println(heros)

	//演示Map的删除
	delete(heros, "hero1")
	fmt.Println(heros)

	//演示Map的查找
	val, ok := cities["no2"]
	if ok {
		fmt.Printf("有no2 key值为%v\n", val)
	} else {
		fmt.Printf("没有no2 key\n")
	}
}
