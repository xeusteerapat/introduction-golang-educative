package main

import "fmt"

func main() {
	myArr := [4]int{0, 1, 2, 5}
	fmt.Print(GetSum(myArr[:]))
}

func GetSum(array []int) int {
	//write code here
	sum := 0

	for _, val := range array {
		sum += val
	}

	return sum //return the sum here
}
