package Q2569

func handleQuery(nums1 []int, nums2 []int, queries [][]int) []int64 {
	var threeCount int
	for _, query := range queries {
		if query[0] == 3 {
			threeCount++
		}
	}
	var originSum int
	var oneCount int
	for i, n := range nums2 {
		originSum += n
		oneCount += nums1[i]
	}
	var sums = make([]int64, 0, threeCount)
	for _, query := range queries {
		switch query[0] {
		case 1:
			for i := query[1]; i <= query[2]; i++ {
				if nums1[i] == 0 {
					nums1[i] = 1
					oneCount++
				} else {
					nums1[i] = 0
					oneCount--
				}
			}
		case 2:
			originSum += oneCount * query[1]
		case 3:
			sums = append(sums, int64(originSum))
		}
	}
	return sums
}
