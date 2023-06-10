package main

import (
	"fmt"

	"githum.com/wujuno/learngo/accounts"
)


func main(){
	account := accounts.NewAccount("junho")
	fmt.Println(account)
}