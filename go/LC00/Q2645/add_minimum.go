package Q2645

func addMinimum(word string) int {
	count := 1
	for i := 1; i < len(word); i++ {
		// ATTENTION: The type of word[i] is byte(uint8).
		// The result will overflow if smaller one subtracts bigger one
		//fmt.Println(word[i] - word[i-1])
		if word[i] <= word[i-1] {
			count++
		}
	}
	return count*3 - len(word)
}
