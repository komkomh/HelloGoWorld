package main

import "fmt"

func main() {

	var xx string = readLine()
	println("hello!" + xx)
}

func readLine() string {
	var str string
	fmt.Scan(&str)
	return str
}
