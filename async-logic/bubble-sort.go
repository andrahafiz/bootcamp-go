package main

import "log"

func main() {
	sorted := BubbleSort(3, 1, 4, 2)
	log.Println("Sorted array", sorted)
}

func BubbleSort(nums ...int) (sorted []int) {

	for i := 0; i < len(nums)-1; i++ {

		for j := i + 1; j < len(nums); j++ {

			log.Println(nums[i], ">", nums[j])

			if nums[i] > nums[j] {
				log.Println("change", nums[i], "in position", i, "to", nums[j], "in position", j)

				temp := nums[i]

				nums[i] = nums[j]

				nums[j] = temp

				log.Println(nums)
			}
		}

		log.Println("hasil sort batch", i+1, ":", nums)

	}
	return
}
