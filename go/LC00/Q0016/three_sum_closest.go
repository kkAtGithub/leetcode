package Q0016

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	var closest = nums[0] + nums[1] + nums[2]
	if closest == target {
		return closest
	}
	n := len(nums)
	sort.Ints(nums)

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, n-1
		for k > j {
			sum := nums[i] + nums[j] + nums[k]
			if abs(sum-target) < abs(closest-target) {
				closest = sum
			}
			if sum > target {
				k--
			} else if sum < target {
				j++
			} else {
				return target
			}
		}

	}

	return closest
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
