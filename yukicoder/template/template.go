package template

import "fmt"

func main() {
	fmt.Printf("delete me")
}

func readString() string {
	var str string
	fmt.Scan(&str)
	return str
}

func readStrings(len int) (strings []string) {
	var str string
	for i := 0; i < len; i++ {
		fmt.Scanf("%s", &str)
		strings = append(strings, str)
	}
	return
}

func readInt() int {
	var num int
	fmt.Scan(&num)
	return num
}

func readNumbers(len int) (nums []int) {
	var num int
	for i := 0; i < len; i++ {
		fmt.Scan(&num)
		nums = append(nums, num)
	}
	return
}
