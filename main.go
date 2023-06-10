package main

import (
	"fmt"

	"github.com/wujuno/learngo/mydict"
)


func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	err :=dictionary.Delete("wf3fdsf")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(baseWord,"is deleted")
	}
}