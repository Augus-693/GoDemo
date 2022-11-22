package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * @Project GoDemo
 * @File 09_strDemo.go
 * @Author Augus Lee
 * @Date 2022/10/7 17:31
 * @Version 1.0
 */

func main() {

	fmt.Println("demo1--------------------------------")
	//统计字符串的长度，按字节 len(str)
	//Golang的编码统一为 utf-8 （ascⅡ 的字符（字母和数字）占一个字节，汉字占三个字节）
	str := "hello编程"
	fmt.Println("str len=", len(str))

	fmt.Println("demo2--------------------------------")
	str2 := "hello编程"
	//字符串遍历，同时处理有中文的问题 r := []rune(str)
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符 = %c\n", r[i])
	}

	fmt.Println("demo3--------------------------------")
	//字符串转整数 n, err := strconv.Atoi(s string)
	n, err := strconv.Atoi("hello")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转换的结果是", n)
	}

	fmt.Println("demo4--------------------------------")
	//整数转字符串 str4 := strconv.Itoa(123456)
	str4 := strconv.Itoa(123456)
	fmt.Printf("str4 = %v, str4 = %T\n", str4, str4)

	fmt.Println("demo5--------------------------------")
	//字符串转 []byte: var bytes = []byte("hello go")
	var bytes = []byte("hello go")
	fmt.Printf("bytes = %v\n", bytes)

	fmt.Println("demo6--------------------------------")
	//[]byte转字符串 ： str = string([]byte(104 101 108 108 111 32 103 111))
	str6 := string([]byte{104, 101, 108, 108, 111, 32, 103, 111})
	fmt.Printf("str6 = %v\n", str6)

	fmt.Println("demo7--------------------------------")
	//十进制转其他进制  : strconv.FormatInt(i int64, base int)
	str7 := strconv.FormatInt(123, 2)
	fmt.Printf("123对应的二进制是%v\n", str7)
	str7 = strconv.FormatInt(123, 16)
	fmt.Printf("123对应的16进制是%v\n", str7)

	fmt.Println("demo8--------------------------------")
	//查找子串是否在指定的字符串中: strings.Contains("seafood", "foo") //true
	b := strings.Contains("seafood", "foo")
	b1 := strings.Contains("seafood", "river")
	fmt.Printf("b = %v\n", b)
	fmt.Printf("b1 = %v\n", b1)

	fmt.Println("demo9--------------------------------")
	//统计一个字符串有几个指定的子串 ： strings.Count("ceheese", "e") //4
	num := strings.Count("succeed", "c")
	fmt.Printf("num = %v\n", num)

	fmt.Println("demo10-------------------------------")
	//不区分大小写的字符串比较(==是区分字母大小写的): fmt.Println(strings.EqualFold("abc", "Abc")) // true
	b = strings.EqualFold("abc", "Abc")
	fmt.Printf("b = %v\n", b)
	fmt.Println("结果", "abc" == "Abc") //==是区分字母大小写的

	fmt.Println("demo11-------------------------------")
	//返回子串在字符串第一次出现的 index 值，如果没有返回-1 : strings.Index("NLT_abc", "abc") // 4
	index := strings.Index("NLT_abc", "abc")
	fmt.Printf("index = %v\n", index)

	fmt.Println("demo12-------------------------------")
	//返回子串在字符串最后一次出现的 index，如没有返回-1 : strings.LastIndex("go golang", "go")
	index = strings.LastIndex("go golang", "go")
	fmt.Printf("index = %v\n", index)

	fmt.Println("demo13-------------------------------")
	//将指定的子串替换成 另外一个子串: strings.Replace("go go hello", "go", "go 语言", n) n 可以指 定你希望替换几个，如果 n=-1 表示全部替换
	str8 := "go go go golang golang"
	str9 := strings.Replace(str8, "go", "hello", -1)
	fmt.Printf("str8 = %v\nstr9 = %v\n", str8, str9)

	fmt.Println("demo14-------------------------------")
	//按照指定的某个字符，为分割标识，将一个字符串拆分成字符串数组：strings.Split("hello,wrold,ok", ",")
	strArr := strings.Split("hello,wrold,ok", ",")
	for i := 0; i < len(strArr); i++ {
		fmt.Printf("str[%v] = %v\n", i, strArr[i])
	}
	fmt.Printf("strArr = %v\n", strArr)

	fmt.Println("demo15-------------------------------")
	//将字符串的字母进行大小写的转换: strings.ToLower("Go") //go strings.ToUpper("Go") // GO
	str10 := "goLang Hello"
	str11 := strings.ToLower(str10)
	str12 := strings.ToUpper(str10)
	fmt.Printf("str11 = %v\n", str11)
	fmt.Printf("str12 = %v\n", str12)

	fmt.Println("demo16-------------------------------")
	//将字符串左右两边的空格去掉： strings.TrimSpace(" tn a lone gopher ntrn   ")
	str = strings.TrimSpace(" tn a lone gopher ntrn   ")
	fmt.Printf("str = %q\n", str)

	fmt.Println("demo17-------------------------------")
	//1将字符串左右两边指定的字符去掉 ： strings.Trim("! hello! ", " !")  // ["hello"] //将左右两边 ! 和 " "去掉
	str = strings.Trim("! he!llo! ", " !")
	fmt.Printf("str = %q\n", str)

	fmt.Println("demo18-------------------------------")
	//将字符串左边指定的字符去掉 ： strings.TrimLeft("! hello! ", " !")  // ["hello"] //将左边 ! 和 " "去掉
	str = strings.TrimLeft("! he!llo! ", " !")
	fmt.Printf("str = %q\n", str)

	fmt.Println("demo19-------------------------------")
	//将字符串右边指定的字符去掉 ：strings.TrimRight("! hello! ", " !")  // ["hello"] //将右边 ! 和 " "去掉
	str = strings.TrimRight("! he!llo! ", " !")
	fmt.Printf("str = %q\n", str)

	fmt.Println("demo20-------------------------------")
	//判断字符串是否以指定的字符串开头: strings.HasPrefix("ftp://47.101.136.27", "ftp") // true
	b = strings.HasPrefix("ftp://47.101.136.27", "ftp")
	fmt.Printf("b = %v\n", b)

	fmt.Println("demo21--------------------------------")
	//判断字符串是否以指定的字符串结束: strings.HasSuffix("NLT_abc.jpg", "abc") //false
	b = strings.HasSuffix("NLT_abc.jpg", "abc")
	fmt.Printf("b = %v\n", b)
}
