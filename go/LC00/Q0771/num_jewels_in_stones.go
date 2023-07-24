package Q0771

func numJewelsInStones(jewels string, stones string) int {
	var jewelSet = make(map[rune]struct{}, len(jewels))
	for _, j := range jewels {
		jewelSet[j] = struct{}{}
	}

	var count int
	for _, s := range stones {
		if _, ok := jewelSet[s]; ok {
			count++
		}
	}

	return count
}
