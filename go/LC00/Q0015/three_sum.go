package Q0015

import "sort"

func threeSum1(nums []int) [][]int {
	sort.Ints(nums)
	if nums[0] > 0 || nums[len(nums)-1] < 0 {
		return nil
	}
	temp := make([]int, 0, len(nums)-2)
	result := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-1; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			target := -nums[i] - nums[j]
			temp = append(temp, nums[j+1:]...)
			if exists(temp, target) {
				result = append(result, []int{nums[i], nums[j], target})
			}
			temp = temp[:0]
		}
	}
	return result
}

func exists(nums []int, target int) bool {
	left, right := 0, len(nums)
	for left < right {
		k := (left + right - 1) / 2
		if nums[k] > target {
			right = k
		} else if nums[k] < target {
			left = k + 1
		} else {
			return true
		}
	}
	return false
}

func threeSum2(nums []int) [][]int {
	sort.Ints(nums)

	var result [][]int
	size := len(nums)
	for i := 0; i < size-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			break
		}
		if nums[i]+nums[size-2]+nums[size-1] < 0 {
			continue
		}
		j, k := i+1, size-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum < 0 {
				j++
			} else if sum > 0 {
				k--
			} else {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			}
		}
	}
	return result
}
