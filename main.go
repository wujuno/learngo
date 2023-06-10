package main

import (
	"fmt"

	"github.com/wujuno/learngo/accounts"
)


func main() {
	account := accounts.NewAccount("junho")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance())
}