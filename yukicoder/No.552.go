package main

import "fmt"

func main() {
	var N string
	fmt.Scan(&N)
	if N == "0" {
		fmt.Println(N)
	} else {
		fmt.Println(N + "0")
	}
}
