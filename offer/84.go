package main

import "sort"

func permuteUnique(nums []int) (ans [][]int) {
	path := make([]int, 0)
	used := make([]bool, len(nums))
	sort.Ints(nums)
	var dfs func()
	dfs = func() {
		if len(path) == len(nums) {
			ans = append(ans, append([]int(nil), path...))
		}
		for i := 0; i < len(nums); i++ {
			if i-1 >= 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			dfs()
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	dfs()
	return
}
