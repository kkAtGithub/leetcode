package _select

import "sort"

func sortSlice(nums sort.IntSlice) {
	for i := 0; i < len(nums); i++ {
		minIdx := i
		for j := i; j < len(nums); j++ {
			if nums[j] < nums[minIdx] {
				minIdx = j
			}
		}
		if minIdx != i {
			nums[i], nums[minIdx] = nums[minIdx], nums[i]
		}
	}
}
