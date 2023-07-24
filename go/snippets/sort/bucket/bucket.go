package bucket

import (
	"math"
	"sort"
)

func sortSlice(nums sort.IntSlice) {
	min, max := math.MaxInt, math.MinInt
	for _, num := range nums {
		if num > max {
			max = num
		} else if num < min {
			min = num
		}
	}
	bucketCount := (max-min)/len(nums) + 1
	buckets := make([]sort.IntSlice, bucketCount)
	for _, num := range nums {
		bucketNo := (num - min) / len(nums)
		buckets[bucketNo] = append(buckets[bucketNo], num)
	}
	i := 0
	for _, bucket := range buckets {
		sort.Sort(bucket)
		for _, b := range bucket {
			nums[i] = b
			i++
		}
	}
}
