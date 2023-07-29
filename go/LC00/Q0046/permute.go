package Q0046

import "math"

var res [][]int

func permute(nums []int) [][]int {
	res = make([][]int, factorial(len(nums)))
	backtrack(nums, make([]int, 0, len(nums)))
	return res
}

func backtrack(nums []int, r []int) {
	if len(r) == len(nums) {
		_r := make([]int, 0, len(r))
		copy(_r, r)
		res = append(res, _r)
		return
	}
	for i, num := range nums {
		if num == math.MinInt {
			continue
		}
		nums[i] = math.MinInt
		r = append(r, num)
		backtrack(nums, r)
		r = r[0 : len(r)-1]
		nums[i] = num
	}
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return factorial(n-1) * n
}
