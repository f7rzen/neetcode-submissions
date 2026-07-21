func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		need := target - num

		if index, ok := seen[need]; ok {
			return []int{index, i}
		}

		seen[num] = i
	}

	return nil
}