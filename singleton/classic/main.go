package main

import "fmt"

func main() {
	singleton := GetInstance()
	fmt.Println(singleton.GetDescription())
}
