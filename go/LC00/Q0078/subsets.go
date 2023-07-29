package Q0078

var res = make([][]int, 0)

func subsets(nums []int) [][]int {
	res = res[:0]
	backtrack(nums, make([]int, 0, len(nums)), 0)
	return res
}

func backtrack(nums, r []int, n int) {
	_r := make([]int, len(r))
	copy(_r, r)
	res = append(res, _r)

	for i := n; i < len(nums); i++ {
		num := nums[i]
		r = append(r, num)
		backtrack(nums, r, i+1)
		r = r[:len(r)-1]
	}
}
