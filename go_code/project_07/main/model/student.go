package model

/**
 * @Project GoDemo
 * @File student.go
 * @Author Augus Lee
 * @Date 2022/10/14 15:14
 * @Version 1.0
 */

type student struct {
	Name  string
	Age   int
	score float64
}

//因为student结构体首字母是小写，因此是只能在model使用
//通过工厂模式来解决

func NewStudent(n string, a int, s float64) *student {
	return &student{
		Name:  n,
		Age:   a,
		score: s,
	}
}

//如果score字段首字母小写，则，在其它包不可以直接方法，我们可以提供一个方法

func (s *student) GetScore() float64 {
	return s.score
}
