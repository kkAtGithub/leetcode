package Q2178

func maximumEvenSplit(finalSum int64) []int64 {
	if finalSum%2 == 1 {
		return nil
	}
	var i int64 = 2
	var ans []int64
	for ; i <= finalSum; i += 2 {
		ans = append(ans, i)
		finalSum -= i
	}
	ans[len(ans)-1] += finalSum
	return ans
}
