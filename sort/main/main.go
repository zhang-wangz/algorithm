package main

import "algorithm-pattern/sort"

func main() {
	nums := []int{9, 2, 5, 4, 8, 2}
	nums = sort.MergeSort(nums)
	for _, val := range nums {
		println(val)
	}
}
