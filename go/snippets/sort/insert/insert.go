package insert

import "sort"

func sortSlice(nums sort.IntSlice) {
	for i := 1; i < len(nums); i++ {
		t := nums[i]
		for j := 0; j < i; j++ {
			if nums[j] <= t {
				continue
			}
			for k := i; k > j; k-- {
				nums[k] = nums[k-1]
			}
			nums[j] = t
			break
		}
	}
}
