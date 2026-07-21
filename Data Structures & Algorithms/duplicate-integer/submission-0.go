func hasDuplicate(nums []int) bool {
	counts := make(map[int]int, len(nums))
	for _, num := range nums {
		_, ok := counts[num]
		if ok {
			counts[num]++
		} else {
			counts[num] = 1
		}
	}

	for _, count := range counts {
		if count > 1 {
			return true
		}
	}
	return false
}
