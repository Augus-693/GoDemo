package main

import "fmt"

/**
 * @Project GoDemo
 * @File 06_mapTraverse.go
 * @Author Augus Lee
 * @Date 2022/10/13 18:39
 * @Version 1.0
 */

func main() {
	//使用for-range遍历Map
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "上海"
	cities["no3"] = "深圳"
	cities["no4"] = "天津"

	for k, v := range cities {
		fmt.Printf("k = %v, v = %v\n", k, v)
	}

	//使用for-range遍历一个结构比较复杂的map（该map 的 value 又是一个map）
	studentMap := make(map[string]map[string]string)

	studentMap["stu01"] = make(map[string]string, 3)
	studentMap["stu01"]["name"] = "jack"
	studentMap["stu01"]["sex"] = "male"
	studentMap["stu01"]["address"] = "北京"

	studentMap["stu02"] = make(map[string]string, 3)
	studentMap["stu02"]["name"] = "mary"
	studentMap["stu02"]["sex"] = "female"
	studentMap["stu02"]["address"] = "上海"

	for k1, v1 := range studentMap {
		fmt.Println("k1 = ", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\t kay = %v value = %v\n", k2, v2)
		}
		fmt.Println("")
	}
}
