func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	i := 0
	j := 0

	result := make([][]int, 0)
	for i < len(firstList) && j < len(secondList) {
		start := max(firstList[i][0], secondList[j][0])
		end := min(firstList[i][1], secondList[j][1])
		if start <= end {
			interval := make([]int, 2)
			interval[0] = start
			interval[1] = end

			result = append(result, interval)
		}

		if firstList[i][1] < secondList[j][1] {
			i++
		} else {
			j++
		}
	}
	return result
}