package Q2544

import "strconv"

func alternateDigitSum(n int) int {
	numStr := strconv.Itoa(n)
	sum := 0
	for i, d := range numStr {
		d := int(d - '0')
		switch i % 2 {
		case 1:
			sum -= d
		case 0:
			sum += d
		}
	}
	return sum
}
