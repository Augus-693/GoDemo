package test

/**
 * @Project GoDemo
 * @File    cal.go
 * @Author  Augus Lee
 * @Date    2022/10/18 17:20
 * @GoVersion go1.19.2
 * @Version 1.0
 */

func addUpper(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}

func getSub(n1 int, n2 int) int {
	return n1 - n2
}
