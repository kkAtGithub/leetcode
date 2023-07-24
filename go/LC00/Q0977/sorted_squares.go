package Q0977

func sortedSquares(nums []int) []int {
	i, j, cursor := 0, len(nums)-1, len(nums)-1
	res := make([]int, len(nums))
	for i <= j && cursor >= 0 {
		if abs(nums[j]) >= abs(nums[i]) {
			res[cursor] = square(nums[j])
			j--
		} else {
			res[cursor] = square(nums[i])
			i++
		}
		cursor--
	}
	return res
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func square(n int) int {
	return n * n
}
