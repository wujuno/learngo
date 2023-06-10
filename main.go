package main

import (
	"fmt"

	"github.com/wujuno/learngo/accounts"
)


func main() {
	account := accounts.NewAccount("junho")
	account.Deposit(10)
	fmt.Println(account)
}