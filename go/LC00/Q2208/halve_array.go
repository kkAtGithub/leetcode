package Q2208

func halveArray(nums []int) int {
	count, sum, length := 0, 0, len(nums)
	for i := 0; i < len(nums); i++ {
		nums[i] <<= 15
		sum += nums[i]
	}

	for i := length/2 - 1; i >= 0; i-- {
		adjustHeap(nums, i, length)
	}

	for sum >>= 1; sum > 0; count++ {
		nums[0] >>= 1
		sum -= nums[0]
		adjustHeap(nums, 0, length)
	}
	return count
}

func adjustHeap(nums []int, i, length int) {
	t := nums[i]
	for k := 2*i + 1; k < length; k = 2*k + 1 {
		if k+1 < length && nums[k+1] > nums[k] {
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
