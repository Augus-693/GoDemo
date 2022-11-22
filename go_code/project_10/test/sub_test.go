package test

import "testing"

/**
 * @Project GoDemo
 * @File    sub_test.go
 * @Author  Augus Lee
 * @Date    2022/10/18 17:21
 * @GoVersion go1.19.2
 * @Version 1.0
 */

func TestGetSub(t *testing.T) {
	res := getSub(10, 5)
	if res != 5 {
		t.Fatalf("getSub(10, 3) 执行错误，期望值=%v 实际值=%v\n", 5, res)
	}

	//如果正确，输出日志
	t.Logf("getSub(10, 3) 执行正确!!!!...")

}
