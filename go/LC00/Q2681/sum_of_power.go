package Q2681

import "sort"

var resPowerOfTwo = make([]int, 0)

func sumOfPower(nums []int) int {
	n := len(nums)
	sort.Ints(nums)

	res := 0
	mod := int(1e9 + 7)
	for i := n - 1; i >= 0; i-- {
		p := power(nums[i], 2) % mod
		for j := i; j >= 0; j-- {
			_p := p * nums[j] % mod
			gap := i - j - 1
			res += _p
			res %= mod
			if gap > 0 {
				times := powerOfTwo(gap)
				for k := 1; k < times; k++ {
					res += _p
					res %= mod
				}
			}
		}
	}
	return res
}

func powerOfTwo(n int) int {
	if n < len(resPowerOfTwo) {
		return resPowerOfTwo[n]
	}
	if n == 0 {
		resPowerOfTwo = append(resPowerOfTwo, 1)
		return 1
	}
	res := powerOfTwo(n-1) * 2
	resPowerOfTwo = append(resPowerOfTwo, res)
	return res
}

func power(n, m int) int {
	res := 1
	for i := 0; i < m; i++ {
		res *= n
	}
	return res
}
