func twoSum(numbers []int, target int) []int {
	start := 0
	end := len(numbers) - 1

	for i := 0; i < len(numbers); i++ {
		if numbers[start]+numbers[end] == target {
			break
		} else {
			if target-numbers[start] < numbers[end] {
				end--
			} else {
				start++
			}
		}
	}
	return []int{start + 1, end+1}
}