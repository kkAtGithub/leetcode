package Q0088

func merge(nums1 []int, m int, nums2 []int, n int) {
	cursor := len(nums1) - 1
	i, j := m-1, n-1
	for i >= 0 && j >= 0 {

		if nums1[i] > nums2[j] {
			nums1[cursor] = nums1[i]
			i--
			cursor--
			continue
		} else if nums1[i] < nums2[j] {
			nums1[cursor] = nums2[j]
			j--
			cursor--
			continue
		} else {
			nums1[cursor] = nums2[j]
			j--
			cursor--
			nums1[cursor] = nums1[i]
			i--
			cursor--
			continue
		}
	}
	for j >= 0 {
		nums1[cursor] = nums2[j]
		j--
		cursor--
	}
}
