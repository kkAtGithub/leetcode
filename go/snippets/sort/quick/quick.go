package quick

import "sort"

func sortSlice(nums sort.IntSlice) {
	_sortSlice(nums, 0, len(nums)-1)
}

func _sortSlice(nums sort.IntSlice, left, right int) {
	if left >= right {
		return
	}

	i, j, pivot := left, right, nums[(left+right)>>1]
	for i < j {
		for nums[i] < pivot {
			i++
		}
		for nums[j] > pivot {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	_sortSlice(nums, left, j)
	_sortSlice(nums, j+1, right)
}
