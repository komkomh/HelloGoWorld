package main

import "fmt"

func main() {
	n, p := 0, 0
	_, _ = fmt.Scan(&n, &p)
	var np = n * p
	if p == np {
		fmt.Print("=")
	} else {
		fmt.Println("!=")
	}
}
