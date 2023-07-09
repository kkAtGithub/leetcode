package Q0004

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n, m := len(nums1), len(nums2)
	l, r := (n+m+1)/2, (n+m+2)/2
	return float64(findKthFromTwoSortedArrays(nums1, 0, n-1, nums2, 0, m-1, l)+findKthFromTwoSortedArrays(nums1, 0, n-1, nums2, 0, m-1, r)) / 2
}

func findKthFromTwoSortedArrays(nums1 []int, start1, end1 int, nums2 []int, start2, end2 int, k int) int {
	len1, len2 := end1-start1+1, end2-start2+1
	if len1 > len2 {
		return findKthFromTwoSortedArrays(nums2, start2, end2, nums1, start1, end1, k)
	}
	if len1 == 0 {
		return nums2[start2+k-1]
	}
	if k == 1 {
		return min(nums2[start2], nums1[start1])
	}
	t1 := start1 + min(len1, k/2) - 1
	t2 := start2 + min(len2, k/2) - 1
	if nums1[t1] > nums2[t2] {
		return findKthFromTwoSortedArrays(nums1, start1, end1, nums2, t2+1, end2, k-(t2-start2+1))
	} else {
		return findKthFromTwoSortedArrays(nums1, t1+1, end1, nums2, start2, end2, k-(t1-start1+1))
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
