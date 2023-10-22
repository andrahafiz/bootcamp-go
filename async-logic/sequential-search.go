package main

import "log"

func main() {
	find := 15

	nums := []int{4, 2, 1, 2, 4, 8, 5, 10, 4, 9, 4}

	found, index := Sequential(find, nums...)
	log.Println(found, index)
}

func Sequential(find int, nums ...int) (found bool, index int) {
	for i, num := range nums {
		if num == find {
			found = true
			index = i
			return
		}
	}

	index = -1
	return
}
