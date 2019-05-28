package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	var V []int = scanNums(N)
	var max int = 0
	for i := 0; i < len(V); i++ {
		var minus2 = 0
		if i-2 >= 0 {
			minus2 = V[i-2]
		}
		var minus3 = 0
		if i-3 >= 0 {
			minus3 = V[i-3]
		}
		if minus2 > minus3 {
			V[i] += minus2
		} else {
			V[i] += minus3
		}
		if V[i] > max {
			max = V[i]
		}
	}
	fmt.Println(max)
}

func scanNums(len int) (nums []int) {
	var num int
	for i := 0; i < len; i++ {
		fmt.Scan(&num)
		nums = append(nums, num)
	}
	return
}
