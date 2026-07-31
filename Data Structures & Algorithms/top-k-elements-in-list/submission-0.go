func topKFrequent(nums []int, k int) []int {

	type Pair struct {
		Num   int
		Count int
	}

	seen := make(map[int]int)
	for _, num := range nums {
		seen[num]++
	}

	pairs := []Pair{}

	for num, count := range seen {
		pairs = append(pairs, Pair{
			Num:   num,
			Count: count,
		})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Count > pairs[j].Count
	})

	result := make([]int, 0, k)
	for i := 0; i < k; i++ {
		result = append(result, pairs[i].Num)
	}

	return result
}