package Q0077

var res = make([][]int, 0)

func combine(n int, k int) [][]int {
	res = res[:0]
	backtrack(n, k, 1, make([]int, 0))
	return res
}

func backtrack(n, k, i int, r []int) {
	if len(r) == k {
		_r := make([]int, len(r))
		copy(_r, r)
		res = append(res, _r)
		return
	}
	for j := i; j <= n; j++ {
		r = append(r, j)
		backtrack(n, k, j+1, r)
		r = r[:len(r)-1]
	}
}
