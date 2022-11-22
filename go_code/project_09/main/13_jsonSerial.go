package main

import (
	"encoding/json"
	"fmt"
)

/**
 * @Project GoDemo
 * @File    13_jsonSerial.go
 * @Author  Augus Lee
 * @Date    2022/10/18 16:08
 * @GoVersion go1.19.2
 * @Version 1.0
 */

// Monster /**
type Monster struct {
	Name     string `json:"monster_name"` //反射机制
	Age      int    `json:"monster_age"`
	Birthday string //....
	Sal      float64
	Skill    string
}

/**
 * @functionName testStruct
 * @author Augus Lee
 * @description 将结构体序列化
 * @date 2022-10-18 16:10:09
 **/
func testStruct() {
	//演示
	monster := Monster{
		Name:     "牛魔王",
		Age:      500,
		Birthday: "2011-11-11",
		Sal:      8000.0,
		Skill:    "牛魔拳",
	}

	//将monster 序列化
	data, err := json.Marshal(&monster) //..
	if err != nil {
		fmt.Printf("序列号错误 err=%v\n", err)
	}
	//输出序列化后的结果
	fmt.Printf("monster序列化后=%v\n", string(data))

}

/**
 * @functionName testMap
 * @author Augus Lee
 * @description 将 Map 序列化
 * @date 2022-10-18 16:10:49
 **/
func testMap() {
	//定义一个map
	var a map[string]interface{}
	//使用map,需要make
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 30
	a["address"] = "洪崖洞"

	//将a这个map进行序列化
	//将monster 序列化
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	//输出序列化后的结果
	fmt.Printf("map 序列化后=%v\n", string(data))

}

/**
 * @functionName testSlice
 * @author Augus Lee
 * @description 将 Slice 序列化
 * @date 2022-10-18 16:11:08
 **/
func testSlice() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	//使用map前，需要先make
	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = "7"
	m1["address"] = "北京"
	slice = append(slice, m1)

	var m2 map[string]interface{}
	//使用map前，需要先make
	m2 = make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = "20"
	m2["address"] = [2]string{"墨西哥", "夏威夷"}
	slice = append(slice, m2)

	//将切片进行序列化操作
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	//输出序列化后的结果
	fmt.Printf("slice 序列化后=%v\n", string(data))

}

// 对基本数据类型序列化，对基本数据类型进行序列化意义不大
func testFloat64() {
	var num1 float64 = 2345.67

	//对num1进行序列化
	data, err := json.Marshal(num1)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	//输出序列化后的结果
	fmt.Printf("num1 序列化后=%v\n", string(data))
}

func main() {
	//演示将结构体, map , 切片进行序列号
	testStruct()
	testMap()
	testSlice()   //演示对切片的序列化
	testFloat64() //演示对基本数据类型的序列化
}
