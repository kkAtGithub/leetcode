package heap

import "sort"

func sortSlice(nums sort.IntSlice) {
	length := len(nums)
	for i := length/2 - 1; i >= 0; i-- {
		adjustHeap(nums, i, length)
	}

	for i := length - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		adjustHeap(nums, 0, i)
	}
}

func adjustHeap(nums sort.IntSlice, i, length int) {
	t := nums[i]
	for k := i*2 + 1; k < length; k = k*2 + 1 {
		if k < length-1 && nums[k+1] > nums[k] {
			k++
		}
		if nums[k] > t {
			nums[i] = nums[k]
			i = k
		} else {
			break
		}
	}
	nums[i] = t
}
