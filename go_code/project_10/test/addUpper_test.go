package test

import "testing"

/**
 * @Project GoDemo
 * @File    addUpper_test.go
 * @Author  Augus Lee
 * @Date    2022/10/18 17:20
 * @GoVersion go1.19.2
 * @Version 1.0
 */

func TestAddUpper(t *testing.T) {
	res := addUpper(10)
	if res == 55 {
		t.Logf("执行正确，期望值=%v 实际值=%v\n", 55, res)
	} else {
		t.Fatalf("执行错误，期望值=%v 实际值=%v\n", 55, res)
	}
}
