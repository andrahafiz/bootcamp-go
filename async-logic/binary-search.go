package main

import (
	"fmt"
	"log"
)

func main() {
	sortedArray := []int{1, 2, 3, 4, 5, 6, 7, 8}
	find := 8
	search := Binary(find, sortedArray...)
	fmt.Println(search)
}

func Binary(find int, sortedNums ...int) (found bool) {
	if len(sortedNums) <= 1 {
		if sortedNums[0] == find {
			return true
		}
		return false
	}

	for {
		middleIndex := len(sortedNums) / 2

		log.Println("middle index", middleIndex, "sorted nums", sortedNums, "target data", find)

		if sortedNums[middleIndex] == find {
			log.Println("done with middle index", middleIndex, "sorted nums", sortedNums, "target data", find)
			return true
		} else if find < sortedNums[middleIndex] {
			sortedNums = sortedNums[:middleIndex]
		} else if find > sortedNums[middleIndex] {
			sortedNums = sortedNums[middleIndex:]
		}

		if middleIndex == 0 {
			break
		}
	}

	return false
}
