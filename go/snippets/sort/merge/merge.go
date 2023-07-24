package merge

import "sort"

func sortSlice(nums sort.IntSlice) {
	temp := make(sort.IntSlice, len(nums))
	_sortSlice(nums, temp, 0, len(nums)-1)
}

func _sortSlice(nums, temp sort.IntSlice, left, right int) {
	if left >= right {
		return
	}
	mid := (left + right) >> 1
	_sortSlice(nums, temp, left, mid)
	_sortSlice(nums, temp, mid+1, right)

	i, j, cursor := left, mid+1, left
	for i <= mid && j <= right && cursor <= right {
		if nums[i] < nums[j] {
			temp[cursor] = nums[i]
			i++
		} else {
			temp[cursor] = nums[j]
			j++
		}
		cursor++
	}
	for cursor <= right {
		if i <= mid {
			temp[cursor] = nums[i]
			i++
		}
		if j <= right {
			temp[cursor] = nums[j]
			j++
		}
		cursor++
	}
	for n := left; n <= right; n++ {
		nums[n] = temp[n]
	}
}
