package main

import "fmt"


func main() {
	junho := map[string]string{"name": "junho", "age": "29"}
	for key, value := range junho {
		fmt.Println(key, value)
	}
}