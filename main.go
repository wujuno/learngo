package main

import (
	"fmt"

	"github.com/wujuno/learngo/accounts"
)


func main(){
	account := accounts.NewAccount("junho")
	fmt.Println(account)
}