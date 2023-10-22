package main

import (
	"fmt"
)

func main() {
	nums := 10
	for i := 1; i <= nums; i++ {
		fmt.Print(fibonaci(i), ",")
	}
}

func fibonaci(nums int) int {
	if nums <= 1 {
		return nums
	}
	return fibonaci(nums-1) + fibonaci(nums-2)
}
