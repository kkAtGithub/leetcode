package Q1911

import "sort"

func maxAlternatingSum1(nums []int) int64 {
	numsWithIndex := make([][2]int, 0, len(nums))
	for i, num := range nums {
		numsWithIndex = append(numsWithIndex, [2]int{
			num,
			i,
		})
	}
	sort.Slice(numsWithIndex, func(i, j int) bool {
		return numsWithIndex[i][0] < numsWithIndex[j][0]
	})

	i, j, sum, lastIndex := 0, len(numsWithIndex)-2, int64(numsWithIndex[len(numsWithIndex)-1][0]), numsWithIndex[len(numsWithIndex)-1][1]
	for i < j {
		if numsWithIndex[i][1] <= lastIndex {
			i++
			continue
		}
		if numsWithIndex[j][1] <= lastIndex {
			j--
			continue
		}
		delta := -int64(numsWithIndex[i][0]) + int64(numsWithIndex[j][0])
		if sum+delta > sum {
			sum += delta
		}
		i++
		j--
	}
	return sum
}

func maxAlternatingSum(nums []int) int64 {
	var d, d1 int64 = 0, 0
	for _, num := range nums {
		n := int64(num)
		d, d1 = max(d, d1+n), max(d1, d-n)
	}
	return max(d1, d)
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
