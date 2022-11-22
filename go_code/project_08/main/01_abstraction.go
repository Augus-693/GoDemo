package main

import (
	"fmt"
	"runtime"
)

/**
 * @Project GoDemo
 * @File 01_abstraction.go
 * @Author Augus Lee
 * @Date 2022/10/14 15:40
 * @Version 1.0
 */

/*
ATM机实现存款取款和查询余额
*/

type Account struct {
	AccountNo int
	Pwd       string
	Balance   float64
}

var i int

// 功能菜单:menu
func (menu Account) Menu() {
	fmt.Println("********************************")
	fmt.Println("********请输入要实现的功能********")
	fmt.Println("************1： 存款*************")
	fmt.Println("************2： 取款*************")
	fmt.Println("************3： 查询余额*********")
	fmt.Println("************4： 退出*************")
}

// 登录:Login
func (login Account) Login(AccountNo int, Pwd string) {
	i++
	if i == 3 {
		fmt.Println("输入错误已达三次，程序将自动退出")
		runtime.Goexit()
	}
	var accountNo int
	var password string
	fmt.Printf("请输入账户和密码:")
	fmt.Scanf("%v %v\n", &accountNo, &password)
	if accountNo != AccountNo && password != Pwd {
		fmt.Println("账户或密码错误，请重新输入")
		login.Login(AccountNo, Pwd)
	} else {
		fmt.Println("登录成功!!!")
	}
}

// 存款服务:Deposite
func (deposite Account) Deposite(Balance float64) float64 {
	balance := 0.0
	fmt.Printf("您现在的余额是：%.2f\n", Balance)
	fmt.Printf("请输入需要存入的金额：")
	fmt.Scanf("%v\n", &balance)
	Balance = balance + Balance
	fmt.Printf("您现在的余额是：%.2f\n", Balance)
	return Balance
}

// 取款服务:WithDraw
func (withdraw Account) Withdraw(Balance float64) float64 {
	balance := 0.0
	fmt.Printf("您现在的余额是：%.2f\n", Balance)
	fmt.Printf("请输入需要取出的金额：")
	fmt.Scanf("%v\n", &balance)
	Balance = Balance - balance
	fmt.Printf("您现在的余额是：%.2f\n", Balance)
	return Balance
}

// 查询服务:Query
func (query Account) Query(Balance float64) {
	fmt.Printf("您现在的余额是：%.2f\n", Balance)
}

// 主函数
func main() {
	var person *Account = &Account{AccountNo: 111, Pwd: "qqq", Balance: 563.74}

	person.Login(person.AccountNo, person.Pwd)
lable:
	person.Menu()
	var choice int
	fmt.Printf("请输入：")
	fmt.Scanf("%v\n", &choice)
	switch choice {
	case 1:
		person.Balance = person.Deposite(person.Balance)
		goto lable
	case 2:
		person.Balance = person.Withdraw(person.Balance)
		goto lable
	case 3:
		person.Query(person.Balance)
		goto lable
	default:
		fmt.Printf("程序将退出,是否确认y/n?")
		var affirm string
		fmt.Scanf("%v\n", &affirm)
		fmt.Printf("%v\n", affirm)
		if affirm == "y" || affirm == "Y" {
			runtime.Goexit()
		} else {
			goto lable
		}
	}
}
