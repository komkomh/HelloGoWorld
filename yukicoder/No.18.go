package main

import "fmt"

func main() {
	var xx string = readLine18()
	println("hello!" + xx)
}

func readLine18() string {
	var str string
	fmt.Scan(&str)
	return str
}
